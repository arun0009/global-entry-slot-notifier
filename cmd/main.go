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

const GlobalEntryUrl = "https://ttp.cbp.dhs.gov/schedulerapi/slots?orderBy=soonest&limit=1&locationId=%s&minimum=1"

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
func appointmentCheck(url string, client HTTPClient, notifier Notifier, topic string, beforeDate time.Time) {
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

	if len(appointments) > 0 {
		appointment := appointments[0]
		appointmentTime, err := time.Parse("2006-01-02T15:04", appointment.StartTimestamp)
		if err != nil {
			log.Printf("Failed to parse appointment time: %v", err)
			return
		}
		if appointmentTime.Before(beforeDate) {
			if err := notifier.Notify(appointment.LocationID, appointment.StartTimestamp, topic); err != nil {
				log.Printf("Failed to send notification: %v", err)
			}
		}
	}
}

func main() {
	var location, notifierType, topic, before string
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

			// Create a closure that captures the arguments and calls appointmentCheck with them.
			appointmentCheckFunc := func() {
				appointmentCheck(url, client, notifier, topic, beforeDate)
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

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
