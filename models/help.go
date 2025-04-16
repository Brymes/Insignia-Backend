package models

// HelpRequest represents a contact/help request in the database
type HelpRequest struct {
	SQLModel
	Name        string `gorm:"" json:"name"`
	PhoneNumber string `gorm:"" json:"phone_number"`
	Email       string `gorm:"" json:"email"`
	Message     string `gorm:"" json:"message"`
	Source      string `gorm:"" json:"source"`
	Status      string `gorm:"" json:"status"`
}
