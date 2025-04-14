package controllers

import (
	"Insignia-Backend/config"
	"Insignia-Backend/models"
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"strings"
	"time"
)

// sendBookingConfirmation sends email notifications to the client and Insignia
// This function has been refactored to accept a more comprehensive booking details struct
func sendBookingConfirmation(booking models.BoilerBooking, clientEmail, firstName, lastName string,
	appointmentDate time.Time, address, postcode string) error {

	// Format date for display
	formattedDate := appointmentDate.Format("Monday, January 2, 2006 at 3:04 PM")

	// Create Mailjet client
	mailjetClient := mailjet.NewMailjetClient(
		config.MailjetAPIKeyPublic,
		config.MailjetAPIKeyPrivate,
	)

	bookingTypeStr := string(booking.BookingType)
	subject := fmt.Sprintf("Your Boiler %s Booking Confirmation", bookingTypeStr)

	// Create detailed HTML content with all form fields
	htmlContent := fmt.Sprintf(`
		<h2>Booking Confirmation</h2>
		<p>Dear %s %s,</p>
		<p>Thank you for booking a boiler %s with Insignia. Your appointment has been scheduled for:</p>
		<p><strong>%s</strong></p>
		
		<h3>Booking Details</h3>
		<table style="border-collapse: collapse; width: 100%%;">
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Address:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s, %s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>User Type:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Property Type:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Bedroom Count:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Bathroom Count:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Boiler Fuel Type:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Boiler Type:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Boiler Age:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Boiler Mounting:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>`,
		firstName, lastName, strings.ToLower(bookingTypeStr), formattedDate,
		address, postcode, booking.UserType, booking.PropertyType,
		booking.BedroomCount, booking.BathroomCount, booking.BoilerFuelType,
		booking.BoilerType, booking.BoilerAge, booking.BoilerMounting)

	// Add conditional fields based on booking type
	if booking.BookingType == models.Installation {
		// Add installation-specific details
		htmlContent += fmt.Sprintf(`
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Boiler Model:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Installation Reason:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>`,
			booking.BoilerModel, booking.Reason)

		if booking.OtherReason != "" {
			htmlContent += fmt.Sprintf(`
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Other Reason:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>`, booking.OtherReason)
		}

		htmlContent += fmt.Sprintf(`
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Express Installation:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>`,
			formatBooleanYesNo(booking.ExpressInstallation))
	} else {
		// Add repair-specific details
		htmlContent += fmt.Sprintf(`
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Issue Type:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>`, booking.Reason)

		if booking.OtherReason != "" {
			htmlContent += fmt.Sprintf(`
			<tr>
				<td style="padding: 8px; border: 1px solid #ddd;"><strong>Other Issue:</strong></td>
				<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
			</tr>`, booking.OtherReason)
		}
	}

	// Close the table and add footer
	htmlContent += fmt.Sprintf(`
		</table>
		
		<p>Our team will contact you before the appointment to confirm details.</p>
		<p>If you need to reschedule or have any questions, please contact us at %s.</p>
		<p>Thank you for choosing Insignia.</p>
	`, config.InsigniaEmail)

	// Create plain text version
	textContent := fmt.Sprintf(
		"Booking Confirmation\n\nDear %s %s,\n\nThank you for booking a boiler %s with Insignia. "+
			"Your appointment has been scheduled for: %s\n\n"+
			"Booking Details:\n"+
			"Address: %s, %s\n"+
			"User Type: %s\n"+
			"Property Type: %s\n"+
			"Bedroom Count: %s\n"+
			"Bathroom Count: %s\n"+
			"Boiler Fuel Type: %s\n"+
			"Boiler Type: %s\n"+
			"Boiler Age: %s\n"+
			"Boiler Mounting: %s\n",
		firstName, lastName, strings.ToLower(bookingTypeStr), formattedDate,
		address, postcode, booking.UserType, booking.PropertyType,
		booking.BedroomCount, booking.BathroomCount, booking.BoilerFuelType,
		booking.BoilerType, booking.BoilerAge, booking.BoilerMounting)

	// Add conditional fields based on booking type for text version
	if booking.BookingType == models.Installation {
		textContent += fmt.Sprintf(
			"Boiler Model: %s\n"+
				"Installation Reason: %s\n",
			booking.BoilerModel, booking.Reason)

		if booking.OtherReason != "" {
			textContent += fmt.Sprintf("Other Reason: %s\n", booking.OtherReason)
		}

		textContent += fmt.Sprintf("Express Installation: %s\n",
			formatBooleanYesNo(booking.ExpressInstallation))
	} else {
		textContent += fmt.Sprintf("Issue Type: %s\n", booking.Reason)
		if booking.OtherReason != "" {
			textContent += fmt.Sprintf("Other Issue: %s\n", booking.OtherReason)
		}
	}

	// Add footer to text version
	textContent += fmt.Sprintf(
		"\nOur team will contact you before the appointment to confirm details.\n\n"+
			"If you need to reschedule or have any questions, please contact us at %s.\n\n"+
			"Thank you for choosing Insignia.",
		config.InsigniaEmail)

	// Create a simplified version for the internal team notification
	internalHtmlContent := fmt.Sprintf(`
		<h3>New Booking Received</h3>
		<p><strong>Customer:</strong> %s %s</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Appointment:</strong> %s</p>
		<p><strong>Address:</strong> %s, %s</p>
		<p><strong>Booking Type:</strong> %s</p>
		<p><strong>Property Type:</strong> %s</p>
		<p><strong>Boiler Details:</strong> %s %s, %s old, %s</p>
		<p><strong>Reason:</strong> %s</p>`,
		firstName, lastName, clientEmail, formattedDate, address, postcode,
		bookingTypeStr, booking.PropertyType, booking.BoilerFuelType,
		booking.BoilerType, booking.BoilerAge, booking.BoilerMounting, booking.Reason)

	if booking.BookingType == models.Installation && booking.ExpressInstallation {
		internalHtmlContent += "<p><strong>Express Installation Requested</strong></p>"
	}

	// Prepare messages
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: config.InsigniaEmail,
				Name:  "Insignia Bookings",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: clientEmail,
					Name:  fmt.Sprintf("%s %s", firstName, lastName),
				},
			},
			Subject:  subject,
			TextPart: textContent,
			HTMLPart: htmlContent,
		},
		{
			From: &mailjet.RecipientV31{
				Email: config.InsigniaEmail,
				Name:  "Insignia Bookings",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: config.InsigniaEmail,
					Name:  "Insignia Team",
				},
			},
			Subject: fmt.Sprintf("New Boiler %s Booking - %s %s", bookingTypeStr, firstName, lastName),
			TextPart: fmt.Sprintf("New booking received:\n\nCustomer: %s %s\nEmail: %s\nAppointment: %s\nAddress: %s, %s\nBooking Type: %s",
				firstName, lastName, clientEmail, formattedDate, address, postcode, bookingTypeStr),
			HTMLPart: internalHtmlContent,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)

	return err
}

// Helper function to format boolean values as Yes/No
func formatBooleanYesNo(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}
