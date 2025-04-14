package controllers

import (
	"Insignia-Backend/models"
	"Insignia-Backend/schemas"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateBoilerInstallation(c *gin.Context) {
	var (
		booking   models.BoilerBooking
		validator schemas.BoilerInstallationPayload
		handler   = schemas.NewHandler("Boiler-Installation", "")
	)

	defer handler.SendResponse(c)

	if cond := schemas.ValidateAndBind(&validator, &booking, c, handler); !cond {
		return
	}

	// Set booking type and organization
	booking.BookingType = models.Installation
	booking.Source = "Web-Booking"

	// Map validated data to model
	booking.UserType = string(validator.UserType)
	booking.PropertyType = string(validator.PropertyType)
	booking.BedroomCount = string(validator.BedroomCount)
	booking.BathroomCount = string(validator.BathroomCount)
	booking.BoilerFuelType = string(validator.BoilerFuelType)
	booking.BoilerType = string(validator.BoilerType)
	booking.BoilerAge = string(validator.BoilerAge)
	booking.BoilerMounting = string(validator.BoilerMounting)
	booking.BoilerModel = validator.BoilerModel
	booking.Reason = validator.InstallationReason
	booking.OtherReason = validator.OtherReason
	booking.ExpressInstallation = validator.ExpressInstallation
	booking.Status = "pending"

	// Add calendar sync functionality
	minDays := 4
	if booking.ExpressInstallation {
		minDays = 2
	}

	// Schedule appointment and get the scheduled date
	appointmentDate, err := scheduleAppointment(booking.BookingType, minDays, validator.FirstName, validator.LastName, validator.Address)
	if err != nil {
		handler.Status, handler.Message = 404, "Failed to schedule appointment in calendar"
		handler.Logger.Panicln(err)
		return
	}

	// Insert booking into database
	handler.Insert(&booking, handler.Logger)

	// Send email notification with all booking details
	err = sendBookingConfirmation(booking, validator.Email, validator.FirstName, validator.LastName,
		appointmentDate, validator.Address, validator.Postcode)
	if err != nil {
		handler.Logger.Println(fmt.Sprintf("Failed to send email notification: %v", err))
	}

	handler.Success, handler.Message = true, "Boiler installation booking created successfully"
}

func CreateBoilerRepair(c *gin.Context) {
	var (
		booking   models.BoilerBooking
		validator schemas.BoilerRepairPayload
		handler   = schemas.NewHandler("Boiler-Repair", "")
	)

	defer handler.SendResponse(c)

	if cond := schemas.ValidateAndBind(&validator, &booking, c, handler); !cond {
		return
	}

	// Set booking type and organization
	booking.BookingType = models.Repair
	booking.Source = "Web-Booking"

	// Map validated data to model
	booking.UserType = string(validator.UserType)
	booking.PropertyType = string(validator.PropertyType)
	booking.BedroomCount = string(validator.BedroomCount)
	booking.BathroomCount = string(validator.BathroomCount)
	booking.BoilerFuelType = string(validator.BoilerFuelType)
	booking.BoilerType = string(validator.BoilerType)
	booking.BoilerAge = string(validator.BoilerAge)
	booking.BoilerMounting = string(validator.BoilerMounting)
	booking.Reason = validator.IssueType
	booking.OtherReason = validator.OtherIssue
	booking.Status = "pending"

	// Add calendar sync functionality - repairs always use standard scheduling (T+4)
	minDays := 4

	// Schedule appointment and get the scheduled date
	appointmentDate, err := scheduleAppointment(booking.BookingType, minDays, validator.FirstName, validator.LastName, validator.Address)
	if err != nil {
		handler.Status, handler.Message = 404, "Failed to schedule appointment in calendar"
		handler.Logger.Panicln(err)
		return
	}

	// Insert booking into database
	handler.Insert(&booking, handler.Logger)

	// Send email notification with all booking details
	err = sendBookingConfirmation(booking, validator.Email, validator.FirstName, validator.LastName,
		appointmentDate, validator.Address, validator.Postcode)
	if err != nil {
		handler.Logger.Println(fmt.Sprintf("Failed to send email notification: %v", err))
	}

	handler.Success, handler.Message = true, "Boiler repair booking created successfully"
}
