package schemas

type BoilerInstallationPayload struct {
	// User Information
	UserType UserType `json:"user_type" binding:"required"`

	// Property Information
	PropertyType  PropertyType  `json:"property_type" binding:"required"`
	BedroomCount  BedroomCount  `json:"bedroom_count" binding:"required"`
	BathroomCount BathroomCount `json:"bathroom_count" binding:"required"`

	// Current Boiler Information
	BoilerFuelType BoilerFuelType `json:"boiler_fuel_type" binding:"required"`
	BoilerType     BoilerType     `json:"boiler_type" binding:"required"`
	BoilerAge      BoilerAge      `json:"boiler_age" binding:"required"`
	BoilerMounting BoilerMounting `json:"boiler_mounting" binding:"required"`
	BoilerModel    string         `json:"boiler_model" binding:"omitempty"`
	// Reason for Installation
	InstallationReason string `json:"installation_reason" binding:"required"`
	OtherReason        string `json:"other_reason" binding:"omitempty"`

	// Express Installation
	ExpressInstallation bool `json:"express_installation" binding:"omitempty"`

	// Contact Information
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Postcode  string `json:"postcode" binding:"required"`

	// Appointment
	PreferredDate string `json:"preferred_date" binding:"required"`
}

type BoilerRepairPayload struct {
	// User Information
	UserType UserType `json:"user_type" binding:"required"`

	// Property Information
	PropertyType  PropertyType  `json:"property_type" binding:"required"`
	BedroomCount  BedroomCount  `json:"bedroom_count" binding:"required"`
	BathroomCount BathroomCount `json:"bathroom_count" binding:"required"`

	// Current Boiler Information
	BoilerFuelType BoilerFuelType `json:"boiler_fuel_type" binding:"required"`
	BoilerType     BoilerType     `json:"boiler_type" binding:"required"`
	BoilerAge      BoilerAge      `json:"boiler_age" binding:"required"`
	BoilerMounting BoilerMounting `json:"boiler_mounting" binding:"required"`

	// Issue Information
	IssueType  string `json:"issue_type" binding:"required"`
	OtherIssue string `json:"other_issue" binding:"omitempty"`

	// Contact Information
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Postcode  string `json:"postcode" binding:"required"`

	// Appointment
	PreferredDate string `json:"preferred_date" binding:"required"`
}
