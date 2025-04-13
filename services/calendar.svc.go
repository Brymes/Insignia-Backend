package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var CalendarSVC CalendarService

func init() {
// TODO initialize CalendarSVC with NewCalendarService
}

// TODO Use Google Calnendar API to create events, macimum of 2 bookings in a day, maximum of 2 events in a day, return error if more than 2 bookings in a day, return error if more than 2 events in a day
// TODO function to check all dates from Today to 30 days from Today, return array of dates with > 2 events in a day

// CalendarService handles Google Calendar operations
type CalendarService struct {
	service    *calendar.Service
	calendarID string // The ID of the calendar to use
}

// NewCalendarService creates a new instance of CalendarService
func NewCalendarService(credentialsJSON []byte, calendarID string) (*CalendarService, error) {
	ctx := context.Background()

	// Configure the Google Calendar API client
	config, err := google.JWTConfigFromJSON(credentialsJSON, calendar.CalendarEventsScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file: %v", err)
	}

	client := config.Client(ctx)

	// Create the Calendar service
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to create Calendar service: %v", err)
	}

	return &CalendarService{
		service:    srv,
		calendarID: calendarID,
	}, nil
}

// CreateBooking creates a new booking event in Google Calendar
// Returns error if there are already 2 bookings on the requested date
func (cs *CalendarService) CreateBooking(summary, description string, startTime, endTime time.Time) error {
	// Check if the date already has 2 or more bookings
	bookingsCount, err := cs.countEventsOnDate(startTime)
	if err != nil {
		return fmt.Errorf("failed to check existing bookings: %v", err)
	}

	if bookingsCount >= 2 {
		return errors.New("maximum number of bookings (2) for this date has been reached")
	}

	// Create the event
	event := &calendar.Event{
		Summary:     summary,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: startTime.Format(time.RFC3339),
			TimeZone: "Europe/London", // Adjust timezone as needed
		},
		End: &calendar.EventDateTime{
			DateTime: endTime.Format(time.RFC3339),
			TimeZone: "Europe/London", // Adjust timezone as needed
		},
	}

	// Insert the event
	_, err = cs.service.Events.Insert(cs.calendarID, event).Do()
	if err != nil {
		return fmt.Errorf("failed to create event: %v", err)
	}

	return nil
}

// countEventsOnDate counts the number of events on a specific date
func (cs *CalendarService) countEventsOnDate(date time.Time) (int, error) {
	// Set time to beginning of day
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	// Set time to end of day
	endOfDay := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())

	// Format times for API
	timeMin := startOfDay.Format(time.RFC3339)
	timeMax := endOfDay.Format(time.RFC3339)

	// Query events for the day
	events, err := cs.service.Events.List(cs.calendarID).
		TimeMin(timeMin).
		TimeMax(timeMax).
		SingleEvents(true).
		Do()

	if err != nil {
		return 0, fmt.Errorf("failed to retrieve events: %v", err)
	}

	return len(events.Items), nil
}

// GetOverbooked returns dates from today to 30 days in the future that have more than 2 events
func (cs *CalendarService) GetOverbooked() ([]time.Time, error) {
	var overbooked []time.Time

	// Get today's date
	today := time.Now()
	// Get date 30 days from today
	thirtyDaysLater := today.AddDate(0, 0, 30)

	// Check each day
	for d := today; d.Before(thirtyDaysLater) || d.Equal(thirtyDaysLater); d = d.AddDate(0, 0, 1) {
		count, err := cs.countEventsOnDate(d)
		if err != nil {
			return nil, fmt.Errorf("failed to check events for %s: %v", d.Format("2006-01-02"), err)
		}

		if count > 2 {
			overbooked = append(overbooked, d)
		}
	}

	return overbooked, nil
}

// IsDateAvailable checks if a specific date has less than 2 bookings
func (cs *CalendarService) IsDateAvailable(date time.Time) (bool, error) {
	count, err := cs.countEventsOnDate(date)
	if err != nil {
		return false, err
	}

	return count < 2, nil
}

// GetAvailableDates returns all available dates within the next 30 days
func (cs *CalendarService) GetAvailableDates() ([]time.Time, error) {
	var availableDates []time.Time

	// Get today's date
	today := time.Now()
	// Get date 30 days from today
	thirtyDaysLater := today.AddDate(0, 0, 30)

	// Check each day
	for d := today; d.Before(thirtyDaysLater) || d.Equal(thirtyDaysLater); d = d.AddDate(0, 0, 1) {
		isAvailable, err := cs.IsDateAvailable(d)
		if err != nil {
			return nil, fmt.Errorf("failed to check availability for %s: %v", d.Format("2006-01-02"), err)
		}

		if isAvailable {
			availableDates = append(availableDates, d)
		}
	}

	return availableDates, nil
}
