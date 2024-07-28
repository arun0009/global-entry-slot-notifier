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
	Topic       string
	ReturnError bool
}

func (m *MockNotifier) Notify(locationID int, topic string) error {
	m.Called = true
	m.LocationID = locationID
	m.Topic = topic
	if m.ReturnError {
		return errors.New("mock error")
	}
	return nil
}

func TestAppointmentCheck(t *testing.T) {
	appointments := []Appointment{
		{LocationID: 123},
	}
	appointmentsJSON, _ := json.Marshal(appointments)

	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(string(appointmentsJSON))),
		},
		Err: nil,
	}
	mockNotifier := &MockNotifier{}

	url := "http://example.com"
	topic := "test-topic"

	appointmentCheck(url, mockClient, mockNotifier, topic)

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
	appointmentsJSON, _ := json.Marshal([]Appointment{})

	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(string(appointmentsJSON))),
		},
		Err: nil,
	}
	mockNotifier := &MockNotifier{}

	url := "http://example.com"
	topic := "test-topic"

	appointmentCheck(url, mockClient, mockNotifier, topic)

	if mockNotifier.Called {
		t.Fatalf("Expected Notify not to be called")
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

	appointmentCheck(url, mockClient, mockNotifier, topic)

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

	appointmentCheck(url, mockClient, mockNotifier, topic)

	if mockNotifier.Called {
		t.Fatalf("Expected Notify not to be called")
	}
}

func TestAppointmentCheckScheduler(t *testing.T) {
	count := 0
	appointmentCheck := func() {
		count++
	}

	go appointmentCheckScheduler(1*time.Second, appointmentCheck)

	time.Sleep(3500 * time.Millisecond)

	if count != 3 {
		t.Errorf("Expected appointmentCheck to be called 3 times, got %d", count)
	}
}
