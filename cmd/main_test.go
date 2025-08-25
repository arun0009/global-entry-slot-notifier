package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

// MockHTTPClient is a mock implementation of the HTTPClient interface for testing.
type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.Response, m.Err
}

func (m *MockHTTPClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return m.Response, m.Err
}

// MockNotifier is a mock implementation of the Notifier interface for testing.
type MockNotifier struct {
	Called      bool
	LocationID  int
	StartTime   string
	Topic       string
	ReturnError bool
}

func (m *MockNotifier) Notify(locationID int, startTime string, topic string) error {
	m.Called = true
	m.LocationID = locationID
	m.StartTime = startTime
	m.Topic = topic
	if m.ReturnError {
		return errors.New("mock error")
	}
	return nil
}

// ---- Helpers for constructing timestamps used by tests ----

const apiTimeLayout = "2006-01-02T15:04"

// makeTS returns an API-formatted timestamp at (now + dayOffset) with H:M.
func makeTS(dayOffset int, hour, min int) string {
	base := time.Now().AddDate(0, 0, dayOffset)
	t := time.Date(base.Year(), base.Month(), base.Day(), hour, min, 0, 0, base.Location())
	return t.Format(apiTimeLayout)
}

// makeJSON is a small helper to marshal appointments.
func makeJSON(appts []Appointment) string {
	b, _ := json.Marshal(appts)
	return string(b)
}

// defaultOpts returns CheckOptions with no time-of-day filtering.
func defaultOpts() CheckOptions {
	return CheckOptions{EarliestMinutes: -1, LatestMinutes: -1}
}

// -----------------------------------------------------------

func TestAppointmentCheck(t *testing.T) {
	futureTime := makeTS(1, 10, 30) // 24h+ from now, API layout "2006-01-02T15:04"
	appointments := []Appointment{
		{LocationID: 123, StartTimestamp: futureTime},
	}
	appointmentsJSON := makeJSON(appointments)

	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(appointmentsJSON)),
		},
		Err: nil,
	}
	mockNotifier := &MockNotifier{}

	url := "http://example.com"
	topic := "test-topic"
	beforeDate := time.Now().Add(48 * time.Hour) // Allow appointments within 48 hours

	appointmentCheck(url, mockClient, mockNotifier, topic, beforeDate, defaultOpts())

	if !mockNotifier.Called {
		t.Fatalf("Expected Notify to be called")
	}
	if mockNotifier.LocationID != 123 {
		t.Errorf("Expected LocationID to be %d, got %d", 123, mockNotifier.LocationID)
	}
	if mockNotifier.Topic != topic {
		t.Errorf("Expected topic to be %s, got %s", topic, mockNotifier.Topic)
	}
}

func TestAppointmentCheck_NoAppointments(t *testing.T) {
	appointmentsJSON := makeJSON([]Appointment{})

	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(appointmentsJSON)),
		},
		Err: nil,
	}
	mockNotifier := &MockNotifier{}

	url := "http://example.com"
	topic := "test-topic"
	beforeDate := time.Now().Add(48 * time.Hour) // Allow appointments within 48 hours

	appointmentCheck(url, mockClient, mockNotifier, topic, beforeDate, defaultOpts())

	if mockNotifier.Called {
		t.Fatalf("Expected Notify not to be called")
	}
}

func TestAppointmentCheck_AppointmentOutsideBeforeDate(t *testing.T) {
	// 3 days from now (outside 48h window), using the API layout so it parses.
	outsideTime := makeTS(3, 9, 0)
	appointments := []Appointment{
		{LocationID: 123, StartTimestamp: outsideTime},
	}
	appointmentsJSON := makeJSON(appointments)

	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(appointmentsJSON)),
		},
		Err: nil,
	}
	mockNotifier := &MockNotifier{}

	url := "http://example.com"
	topic := "test-topic"
	beforeDate := time.Now().Add(48 * time.Hour) // Only allow appointments within 2 days

	appointmentCheck(url, mockClient, mockNotifier, topic, beforeDate, defaultOpts())

	if mockNotifier.Called {
		t.Fatalf("Expected Notify not to be called for appointments after beforeDate")
	}
}

func TestAppointmentCheck_HTTPError(t *testing.T) {
	mockClient := &MockHTTPClient{
		Response: nil,
		Err:      errors.New("http error"),
	}
	mockNotifier := &MockNotifier{}

	url := "http://example.com"
	topic := "test-topic"
	beforeDate := time.Now().Add(48 * time.Hour) // Allow appointments within 48 hours

	appointmentCheck(url, mockClient, mockNotifier, topic, beforeDate, defaultOpts())

	if mockNotifier.Called {
		t.Fatalf("Expected Notify not to be called")
	}
}

func TestAppointmentCheck_UnmarshalError(t *testing.T) {
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("invalid json")),
		},
		Err: nil,
	}
	mockNotifier := &MockNotifier{}

	url := "http://example.com"
	topic := "test-topic"
	beforeDate := time.Now().Add(48 * time.Hour) // Allow appointments within 48 hours

	appointmentCheck(url, mockClient, mockNotifier, topic, beforeDate, defaultOpts())

	if mockNotifier.Called {
		t.Fatalf("Expected Notify not to be called")
	}
}

func TestAppointmentCheckScheduler(t *testing.T) {
	count := 0
	appointmentCheckFn := func() {
		count++
	}

	go appointmentCheckScheduler(1*time.Second, appointmentCheckFn)

	time.Sleep(3500 * time.Millisecond)

	if count != 3 {
		t.Errorf("Expected appointmentCheck to be called 3 times, got %d", count)
	}
}

// ---------------- New tests for earliest/latest filtering ----------------

func TestAppointmentCheck_EarliestOnly(t *testing.T) {
	// earliest = 08:00; appointments at 07:00 (filtered) and 09:00 (valid)
	appts := []Appointment{
		{LocationID: 1, StartTimestamp: makeTS(1, 7, 0)},
		{LocationID: 2, StartTimestamp: makeTS(1, 9, 0)},
	}
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(makeJSON(appts))),
		},
	}
	mockNotifier := &MockNotifier{}
	url := "http://example.com"
	topic := "topic"
	beforeDate := time.Now().Add(72 * time.Hour)

	opts := CheckOptions{
		EarliestMinutes: 8 * 60,
		LatestMinutes:   -1,
	}

	appointmentCheck(url, mockClient, mockNotifier, topic, beforeDate, opts)

	if !mockNotifier.Called {
		t.Fatalf("Expected Notify to be called for appointment >= earliest")
	}
	// Should notify on the first valid (09:00), not the 07:00 one.
	if mockNotifier.LocationID != 2 {
		t.Errorf("Expected LocationID 2 (09:00), got %d", mockNotifier.LocationID)
	}
}

func TestAppointmentCheck_LatestOnly(t *testing.T) {
	// latest = 10:00; appointments at 09:30 (valid) then 11:00 (filtered)
	appts := []Appointment{
		{LocationID: 10, StartTimestamp: makeTS(1, 9, 30)},
		{LocationID: 11, StartTimestamp: makeTS(1, 11, 0)},
	}
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(makeJSON(appts))),
		},
	}
	mockNotifier := &MockNotifier{}
	url := "http://example.com"
	topic := "topic"
	beforeDate := time.Now().Add(72 * time.Hour)

	opts := CheckOptions{
		EarliestMinutes: -1,
		LatestMinutes:   10 * 60,
	}

	appointmentCheck(url, mockClient, mockNotifier, topic, beforeDate, opts)

	if !mockNotifier.Called {
		t.Fatalf("Expected Notify to be called for appointment <= latest")
	}
	if mockNotifier.LocationID != 10 {
		t.Errorf("Expected LocationID 10 (09:30), got %d", mockNotifier.LocationID)
	}
}

func TestAppointmentCheck_BothValidRange(t *testing.T) {
	// earliest=08:00, latest=10:00; appointments at 07:30 (filtered), 09:00 (valid), 10:30 (filtered)
	appts := []Appointment{
		{LocationID: 101, StartTimestamp: makeTS(1, 7, 30)},
		{LocationID: 102, StartTimestamp: makeTS(1, 9, 0)},
		{LocationID: 103, StartTimestamp: makeTS(1, 10, 30)},
	}
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(makeJSON(appts))),
		},
	}
	mockNotifier := &MockNotifier{}
	url := "http://example.com"
	topic := "topic"
	beforeDate := time.Now().Add(72 * time.Hour)

	opts := CheckOptions{
		EarliestMinutes: 8 * 60,
		LatestMinutes:   10 * 60,
	}

	appointmentCheck(url, mockClient, mockNotifier, topic, beforeDate, opts)

	if !mockNotifier.Called {
		t.Fatalf("Expected Notify to be called for appointment within [earliest, latest]")
	}
	if mockNotifier.LocationID != 102 {
		t.Errorf("Expected LocationID 102 (09:00), got %d", mockNotifier.LocationID)
	}
}

func TestAppointmentCheck_BothInvalidRange_NoNotify(t *testing.T) {
	// earliest=12:00, latest=10:00, impossible to satisfy both; expect no notify
	appts := []Appointment{
		{LocationID: 201, StartTimestamp: makeTS(1, 9, 0)},
		{LocationID: 202, StartTimestamp: makeTS(1, 11, 30)},
		{LocationID: 203, StartTimestamp: makeTS(1, 13, 0)},
	}
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(makeJSON(appts))),
		},
	}
	mockNotifier := &MockNotifier{}
	url := "http://example.com"
	topic := "topic"
	beforeDate := time.Now().Add(72 * time.Hour)

	opts := CheckOptions{
		EarliestMinutes: 12 * 60,
		LatestMinutes:   10 * 60,
	}

	appointmentCheck(url, mockClient, mockNotifier, topic, beforeDate, opts)

	if mockNotifier.Called {
		t.Fatalf("Expected Notify NOT to be called when latest < earliest")
	}
}