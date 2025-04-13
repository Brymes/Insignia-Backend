package models

import (
	"github.com/google/uuid"
)

type BookingType string

const (
	Installation BookingType = "Installation"
	Repair       BookingType = "Repair"
	Service      BookingType = "Service"
)

type BoilerBooking struct {
	SQLModel
	OrganizationId uuid.UUID `gorm:"not null" json:"organization_id"`

	// Booking Type
	BookingType BookingType `gorm:"not null" json:"booking_type"`

	// User Information
	UserType string `gorm:"not null" json:"user_type"`

	// Property Information
	PropertyType  string `gorm:"not null" json:"property_type"`
	BedroomCount  string `gorm:"not null" json:"bedroom_count"`
	BathroomCount string `gorm:"not null" json:"bathroom_count"`

	// Current Boiler Information
	BoilerFuelType string `gorm:"not null" json:"boiler_fuel_type"`
	BoilerType     string `gorm:"not null" json:"boiler_type"`
	BoilerAge      string `gorm:"not null" json:"boiler_age"`
	BoilerMounting string `gorm:"not null" json:"boiler_mounting"`
	BoilerModel    string `gorm:""         json:"boiler_model"`
	// Reason for Installation/Issue for Repair
	Reason      string `gorm:"not null" json:"reason"`
	OtherReason string `gorm:""         json:"other_reason"`

	// Express Installation
	ExpressInstallation bool `gorm:""         json:"express_installation"`

	// Contact Information
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null" json:"email"`
	Phone     string `gorm:"not null" json:"phone"`
	Address   string `gorm:"not null" json:"address"`
	Postcode  string `gorm:"not null" json:"postcode"`

	// Appointment
	PreferredDate string `gorm:"not null" json:"preferred_date"`

	// Status
	Status string `gorm:"not null;default:'pending'" json:"status"`

	// Source tracking
	Source string `gorm:""         json:"source"`
}
