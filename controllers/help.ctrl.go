package controllers

import (
	"Insignia-Backend/models"
	"Insignia-Backend/schemas"
	"Insignia-Backend/config"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mailjet/mailjet-apiv3-go/v4"
)

func CreateHelpRequest(c *gin.Context) {
	var (
		helpRequest models.HelpRequest
		validator   schemas.HelpRequestPayload
		handler     = schemas.NewHandler("Help-Request", "")
	)

	defer handler.SendResponse(c)

	if cond := schemas.ValidateAndBind(&validator, &helpRequest, c, handler); !cond {
		return
	}

	// Set request source
	helpRequest.Source = "Web-Contact"
	helpRequest.Status = "pending"

	// Map validated data to model
	helpRequest.Name = validator.Name
	helpRequest.PhoneNumber = validator.PhoneNumber
	helpRequest.Email = validator.Email
	helpRequest.Message = validator.Message

	// Insert help request into database
	handler.Insert(&helpRequest, handler.Logger)

	// Send email notification to admin
	err := sendHelpRequestNotification(helpRequest)
	if err != nil {
		handler.Logger.Println(fmt.Sprintf("Failed to send help request notification: %v", err))
	}

	handler.Success, handler.Message = true, "Help request submitted successfully"
}

// sendHelpRequestNotification sends an email notification to the admin about the new help request
func sendHelpRequestNotification(request models.HelpRequest) error {
	// Create Mailjet client
	mailjetClient := mailjet.NewMailjetClient(config.MailjetAPIKeyPublic, config.MailjetAPIKeyPrivate)

	subject := "New Help Request Received"

	// Create HTML content for the admin notification
	htmlContent := fmt.Sprintf(`
		<h2>New Help Request</h2>
		<p>A new help request has been submitted through the contact form.</p>
		
		<h3>Request Details</h3>
		<table style="border-collapse: collapse; width: 100%%;">
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Name:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Phone Number:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>`,
		request.Name, request.PhoneNumber)

	// Add email if provided
	if request.Email != "" {
		htmlContent += fmt.Sprintf(`
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Email:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>`, request.Email)
	}

	// Add message if provided
	if request.Message != "" {
		htmlContent += fmt.Sprintf(`
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Message:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>`, request.Message)
	}

	// Add submission time
	htmlContent += fmt.Sprintf(`
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Submitted:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
		</table>
		
		<p>Please respond to this request as soon as possible.</p>`,
		time.Now().Format("Monday, January 2, 2006 at 3:04 PM"))

	// Create plain text version
	textContent := fmt.Sprintf(
		"New Help Request\n\n"+
			"A new help request has been submitted through the contact form.\n\n"+
			"Request Details:\n"+
			"Name: %s\n"+
			"Phone Number: %s\n",
		request.Name, request.PhoneNumber)

	// Add email and message to text version if provided
	if request.Email != "" {
		textContent += fmt.Sprintf("Email: %s\n", request.Email)
	}
	if request.Message != "" {
		textContent += fmt.Sprintf("Message: %s\n", request.Message)
	}

	// Add submission time to text version
	textContent += fmt.Sprintf(
		"Submitted: %s\n\n"+
			"Please respond to this request as soon as possible.",
		time.Now().Format("Monday, January 2, 2006 at 3:04 PM"))

	// Prepare confirmation email for the user if they provided an email
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: config.InsigniaEmail,
				Name:  "Insignia Support",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: config.InsigniaEmail,
					Name:  "Insignia Team",
				},
			},
			Subject:  subject,
			TextPart: textContent,
			HTMLPart: htmlContent,
		},
	}

	// Add confirmation email to the user if they provided an email
	if request.Email != "" {
		userHtmlContent := fmt.Sprintf(`
			<h2>Thank You for Contacting Us</h2>
			<p>Dear %s,</p>
			<p>Thank you for reaching out to Insignia. We have received your request and our team will contact you shortly.</p>
			<p>For urgent matters, please call us directly at %s.</p>
			<p>Thank you for choosing Insignia.</p>
		`, request.Name, config.InsigniaPhone)

		userTextContent := fmt.Sprintf(
			"Thank You for Contacting Us\n\n"+
				"Dear %s,\n\n"+
				"Thank you for reaching out to Insignia. We have received your request and our team will contact you shortly.\n\n"+
				"For urgent matters, please call us directly at %s.\n\n"+
				"Thank you for choosing Insignia.",
			request.Name, config.InsigniaPhone)

		messagesInfo = append(messagesInfo, mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: config.InsigniaEmail,
				Name:  "Insignia Support",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: request.Email,
					Name:  request.Name,
				},
			},
			Subject:  "We've Received Your Request - Insignia",
			TextPart: userTextContent,
			HTMLPart: userHtmlContent,
		})
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)

	return err
}