package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

const hhmmLayout = "15:04" // Go’s layout for HH:MM (24-hour)

// Set to 10 to search through multiple time slots.
const GlobalEntryUrl = "https://ttp.cbp.dhs.gov/schedulerapi/slots?orderBy=soonest&limit=10&locationId=%s&minimum=1"

// Appointment represents the structure of the appointment JSON returned by the Global Entry API response.
type Appointment struct {
	LocationID     int    `json:"locationId"`
	StartTimestamp string `json:"startTimestamp"`
	EndTimestamp   string `json:"endTimestamp"`
	Active         bool   `json:"active"`
	Duration       int    `json:"duration"`
	RemoteInt      bool   `json:"remoteInd"`
}

// HTTPClient is an interface for making HTTP requests.
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

// Notifier is an interface for sending notifications.
type Notifier interface {
	Notify(locationID int, startTime string, topic string) error
}

// CheckOptions holds optional time-of-day filters in minutes after midnight; -1 means unset.
type CheckOptions struct {
	EarliestMinutes int
	LatestMinutes   int
}

// AppNotifier sends notifications via an app.
type AppNotifier struct {
	Client HTTPClient
}

func (a AppNotifier) Notify(locationID int, startTime string, topic string) error {
	_, err := a.Client.Post(fmt.Sprintf("https://ntfy.sh/%s", topic), "text/plain",
		strings.NewReader(fmt.Sprintf("Global Entry appointment available at %d on %s", locationID, startTime)))
	return err
}

// SystemNotifier sends system notifications.
type SystemNotifier struct{}

func (s SystemNotifier) Notify(locationID int, startTime string, topic string) error {
	return beeep.Notify("Appointment Slot Available", fmt.Sprintf("Appointment at %d on %s", locationID, startTime), "assets/information.png")
}

// parseHHMM parses "HH:MM" in 24-hour format and returns minutes since midnight.
// If input is invalid, returns -1.
func parseHHMM(s string) int {
	if s == "" {
		return -1
	}
	t, err := time.Parse(hhmmLayout, s)
	if err != nil {
		return -1
	}
	return t.Hour()*60 + t.Minute()
}

// appointmentCheckScheduler calls the provided appointmentCheck function at regular intervals.
func appointmentCheckScheduler(interval time.Duration, appointmentCheck func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			appointmentCheck()
		}
	}
}

// appointmentCheck retrieves the appointment slots and triggers the appropriate notifier.
func appointmentCheck(url string, client HTTPClient, notifier Notifier, topic string, beforeDate time.Time, opts CheckOptions) {
	response, err := client.Get(url)
	if err != nil {
		log.Printf("Failed to get appointment slots: %v", err)
		return
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return
	}

	var appointments []Appointment
	err = json.Unmarshal(responseData, &appointments)
	if err != nil {
		log.Printf("Failed to unmarshal response data: %v", err)
		return
	}

	found := false

	for _, appointment := range appointments {
		appointmentTime, err := time.Parse("2006-01-02T15:04", appointment.StartTimestamp)
		if err != nil {
			log.Printf("Failed to parse appointment time: %v", err)
			continue
		}

		// Filter out appointments after cutoff date
		if appointmentTime.After(beforeDate) {
			continue
		}

		// Optional time-of-day window
		if opts.EarliestMinutes >= 0 || opts.LatestMinutes >= 0 {
			minOfDay := appointmentTime.Hour()*60 + appointmentTime.Minute()
			if opts.EarliestMinutes >= 0 && minOfDay < opts.EarliestMinutes {
				continue
			}
			if opts.LatestMinutes >= 0 && minOfDay > opts.LatestMinutes {
				continue
			}
		}

		// Valid appointment
		if err := notifier.Notify(appointment.LocationID, appointment.StartTimestamp, topic); err != nil {
			log.Printf("Failed to send notification for %s: %v", appointment.StartTimestamp, err)
			continue
		}
		found = true
		break // notify once per cycle (first valid, API is orderBy=soonest)
	}

	if !found {
		log.Printf("[%s] No valid appointments found", time.Now().Format("2006-01-02 15:04:05"))
	}
}

func main() {
	var location, notifierType, topic, before string
	var earliestStr, latestStr string
	var interval time.Duration

	rootCmd := &cobra.Command{
		Use:   "global-entry-slot-notifier",
		Short: "Checks for appointment slots and sends notifications",
		Run: func(cmd *cobra.Command, args []string) {
			// Function to prompt user for input
			getInput := func(prompt string) string {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print(prompt)
				input, _ := reader.ReadString('\n')
				return strings.TrimSpace(input)
			}

			// Check and prompt for missing flags
			if location == "" {
				location = getInput("Enter the location ID: ")
			}
			if notifierType == "" {
				notifierType = getInput("Enter the notifier type (app/system): ")
			}
			if notifierType == "app" && topic == "" {
				topic = getInput("Enter the ntfy topic: ")
			}

			// Validate flags
			if location == "" || notifierType == "" || (notifierType == "app" && topic == "") {
				fmt.Println("Both --location and --notifier flags are required. If notifier is app, --topic is required.")
				_ = cmd.Usage()
				os.Exit(1)
			}

			url := fmt.Sprintf(GlobalEntryUrl, location)

			var notifier Notifier
			client := &http.Client{}
			switch notifierType {
			case "app":
				notifier = AppNotifier{Client: client}
			case "system":
				notifier = SystemNotifier{}
			default:
				log.Fatalf("Unknown notifier type: %s", notifierType)
			}

			beforeDate := time.Now().AddDate(1, 0, 0) // Default: 1 year from now
			if before != "" {
				parsedBefore, err := time.Parse("2006-01-02", before)
				if err == nil {
					beforeDate = parsedBefore
				} else {
					log.Printf("Invalid before date format, using default (1 year from now)")
				}
			}

			// Build options from optional earliest/latest (HH:MM, 24-hour). If invalid, ignore.
			earliestMin := parseHHMM(earliestStr)
			latestMin := parseHHMM(latestStr)

			if earliestStr != "" && earliestMin < 0 {
				log.Printf("Invalid --earliest (expected HH:MM), ignoring")
				earliestMin = -1
			}
			if latestStr != "" && latestMin < 0 {
				log.Printf("Invalid --latest (expected HH:MM), ignoring")
				latestMin = -1
			}
			if earliestMin >= 0 && latestMin >= 0 && latestMin < earliestMin {
				log.Printf("--latest is before --earliest; ignoring time window")
				earliestMin, latestMin = -1, -1
			}

			opts := CheckOptions{
				EarliestMinutes: earliestMin,
				LatestMinutes:   latestMin,
			}

			// Create a closure that captures the arguments and calls appointmentCheck with them.
			appointmentCheckFunc := func() {
				appointmentCheck(url, client, notifier, topic, beforeDate, opts)
			}

			go appointmentCheckScheduler(interval, appointmentCheckFunc)

			// Keep the main function running to allow the ticker to tick.
			select {}
		},
	}

	rootCmd.Flags().StringVarP(&location, "location", "l", "", "Specify the location ID")
	rootCmd.Flags().StringVarP(&notifierType, "notifier", "n", "", "Specify the notifier type (app or system)")
	rootCmd.Flags().StringVarP(&topic, "topic", "t", "", "Specify the ntfy topic (required if notifier is app)")
	rootCmd.Flags().DurationVarP(&interval, "interval", "i", time.Second*60, "Specify the interval")
	rootCmd.Flags().StringVarP(&before, "before", "b", "", "Show only appointments before the specified date (YYYY-MM-DD)")
	rootCmd.Flags().StringVarP(&earliestStr, "earliest", "e", "", "Only appointments at/after this time (HH:MM, 24-hour)")
	rootCmd.Flags().StringVarP(&latestStr, "latest", "L", "", "Only appointments at/before this time (HH:MM, 24-hour)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
