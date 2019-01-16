package main

import (
	"bytes"
	"encoding/json"
	"msgbot/models/messenger"
	"os"
)

// ProcessMessage -
// Processes the message and sends a message
// based on the received message
func ProcessMessage(event models.Messaging) {

	userID := event.Sender.ID
	text := event.Message.Text
	var body *bytes.Buffer

	switch text {
	case "Hello":
		user := GetUserInfo(userID)
		respText := "Hello, " + user.FirstName
		body = createTextMessage(userID, respText)
	case "Play":
		respText := "Sorry, not implemented yet."
		body = createTextMessage(userID, respText)
	default:
		respText := "I can't respond that."
		body = createTextMessage(userID, respText)
	}

	url := os.Getenv("POST_MESSAGE")
	PostMessage(url, body)
}

func createButtonMessage(userID string, text string, buttonText string) *bytes.Buffer {
	var response models.Response
	response = models.Response{
		Recipient: models.User{
			ID: userID,
		},
		Message: models.Message{
			Attachment: &models.Attachment{
				Type: "template",
				Payload: models.Payload{
					TemplateType: "button",
					Text:         text,
					Buttons: []models.Button{
						models.Button{
							Type:    "postback",
							Title:   buttonText,
							Payload: "created",
						},
					},
				},
			},
		},
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)

	return body
}

func createTextMessage(userID string, text string) *bytes.Buffer {
	var response models.Response
	response = models.Response{
		Recipient: models.User{
			ID: userID,
		},
		Message: models.Message{
			Text: text,
		},
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)

	return body
}

func createImageMessage(userID string, image string) *bytes.Buffer {
	var response models.Response
	response = models.Response{
		Recipient: models.User{
			ID: userID,
		},
		Message: models.Message{
			Attachment: &models.Attachment{
				Type: "image",
				Payload: models.Payload{
					URL: image,
				},
			},
		},
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)

	return body
}
