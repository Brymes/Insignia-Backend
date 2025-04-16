package schemas

// HelpRequestPayload defines the structure for incoming help request data
type HelpRequestPayload struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"omitempty,email"`
	Message     string `json:"message" binding:"omitempty"`
}