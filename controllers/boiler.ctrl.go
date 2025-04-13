package controllers

import (
	"Insignia-Backend/models"
	"Insignia-Backend/schemas"

	"github.com/gin-gonic/gin"
)

// TODO after Insert on BoilerInstallation or BoilerRepir, add event to calendar on an available date and send mail notifications using mailjet to client email and config.InsigniaEmail

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

	// Insert booking into database
	handler.Insert(&booking, handler.Logger)

	// TODO: Add calendar sync functionality here

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

	// Insert booking into database
	handler.Insert(&booking, handler.Logger)

	// TODO: Add calendar sync functionality here

	handler.Success, handler.Message = true, "Boiler repair booking created successfully"
}

// func GetOrganizationBookings(c *gin.Context) {
// 	var (
// 		bookings []models.BoilerBooking
// 		handler  = schemas.NewHandler("Get-Organization-Bookings", "")
// 		user     = authMacro(c, models.SoroAdmin, handler)
// 	)

// 	defer handler.SendResponse(c)

// 	handler.FetchByOrganizationID(user.OrganizationID, &bookings, handler.Logger)
// 	handler.Success, handler.Message, handler.ResponseData = true, "Bookings fetched successfully", bookings
// }
