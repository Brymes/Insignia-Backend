package controllers

import (
	"Insignia-Backend/models"
	"Insignia-Backend/services"
	"fmt"
	"log"
	"os"
	"time"
)


var CalendarSVC services.CalendarService

func init() {
	// Read credentials file
	credentialsFile := "credentials.json"
	credentialsData, err := os.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	// Initialize calendar service with your calendar ID
	// Replace "primary" with your specific calendar ID if not using the primary calendar
	calendarService, err := services.NewCalendarService(credentialsData, "primary")
	if err != nil {
		log.Fatalf("Failed to initialize calendar service: %v", err)
	}

	// Assign to global variable
	CalendarSVC = *calendarService
}

// scheduleAppointment finds an available date and adds it to the calendar
func scheduleAppointment(bookingType models.BookingType, minDays int, firstName, lastName, address string) (time.Time, error) {
	// Get all available dates for the next 30 days
	availableDates, err := CalendarSVC.GetAvailableDates()
	if err != nil {
		return time.Time{}, fmt.Errorf("error fetching available dates: %v", err)
	}

	// Calculate the minimum date based on minDays
	minDate := time.Now().AddDate(0, 0, minDays)
	
	// Filter dates that are after minDate and not on weekends
	var eligibleDates []time.Time
	for _, date := range availableDates {
		// Skip dates before minDate
		if date.Before(minDate) {
			continue
		}
		
		// Skip weekends
		if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			continue
		}
		
		eligibleDates = append(eligibleDates, date)
	}
	
	// Check if we have any eligible dates
	if len(eligibleDates) == 0 {
		return time.Time{}, fmt.Errorf("no available dates found within the next 30 days")
	}
	
	// Use the first eligible date
	selectedDate := eligibleDates[0]
	
	// Set appointment time to 9:00 AM
	appointmentTime := time.Date(
		selectedDate.Year(), selectedDate.Month(), selectedDate.Day(),
		9, 0, 0, 0, selectedDate.Location(),
	)
	
	// Create a 3-hour appointment slot
	endTime := appointmentTime.Add(3 * time.Hour)
	
	// Create event summary and description
	summary := fmt.Sprintf("%s Booking - %s %s", bookingType, firstName, lastName)
	description := fmt.Sprintf("Boiler %s at %s\nCustomer: %s %s",
		bookingType, address, firstName, lastName)
	
	// Add event to calendar
	err = CalendarSVC.CreateBooking(summary, description, appointmentTime, endTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("error creating calendar event: %v", err)
	}
	
	return appointmentTime, nil
}