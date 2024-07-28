package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
	Notify(locationID int, topic string) error
}

// AppNotifier sends notifications via an app.
type AppNotifier struct {
	Client HTTPClient
}

func (a AppNotifier) Notify(locationID int, topic string) error {
	_, err := a.Client.Post(fmt.Sprintf("https://ntfy.sh/%s", topic), "text/plain",
		strings.NewReader(fmt.Sprintf("There is a global entry appointment open at %d", locationID)))
	return err
}

// SystemNotifier sends system notifications.
type SystemNotifier struct{}

func (s SystemNotifier) Notify(locationID int, topic string) error {
	return beeep.Notify("Appointment Slot Available", fmt.Sprintf("There is a global entry appointment open at %d", locationID), "assets/information.png")
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
func appointmentCheck(url string, client HTTPClient, notifier Notifier, topic string) {
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
		locationID := appointments[0].LocationID
		if err := notifier.Notify(locationID, topic); err != nil {
			log.Printf("Failed to send notification: %v", err)
		}
	}
}

func main() {
	// Define command-line flags
	location := flag.String("location", "", "Specify the location ID")
	notifierType := flag.String("notifier", "", "Specify the notifier type (app/system)")
	topic := flag.String("topic", "", "Specify the ntfy topic (required if notifier is app)")
	interval := flag.Duration("interval", time.Second*60, "Specify the interval")

	// Parse command-line flags
	flag.Parse()

	// Validate flags
	if *location == "" || *notifierType == "" || (*notifierType == "app" && *topic == "") {
		fmt.Println("Both --location and --notifier flags are required. If notifier is app, --topic is required.")
		flag.Usage()
		os.Exit(1)
	}

	url := fmt.Sprintf(GlobalEntryUrl, *location)

	var notifier Notifier
	client := &http.Client{}
	switch *notifierType {
	case "app":
		notifier = AppNotifier{Client: client}
	case "system":
		notifier = SystemNotifier{}
	default:
		log.Fatalf("Unknown notifier type: %s", *notifierType)
	}

	// Create a closure that captures the arguments and calls appointmentCheck with them.
	appointmentCheckFunc := func() {
		appointmentCheck(url, client, notifier, *topic)
	}

	go appointmentCheckScheduler(*interval, appointmentCheckFunc)

	// Keep the main function running to allow the ticker to tick.
	select {}
}
