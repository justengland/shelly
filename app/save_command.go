// dump_command.go
package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strings"
)

const FileNamePrompt = "Generate a valid yaml file name based on the following json messages, only return the filename no description\n---\n"

func NewDumpCommand() CommandHandler {
	return CommandHandler{
		Name:    "save",
		Command: saveCommand,
	}
}

func saveCommand(userMessage string) string {
	log.Infof("save: " + userMessage)
	var fileName string
	if userMessage == "/save" {
		// todo call gpt to get the file name
		jsonBytes, err := json.Marshal(chat)
		if err != nil {
			log.Fatalf("Failed to marshal object to JSON: %v", err)
		}

		// Convert the JSON byte array to a string
		jsonString := string(jsonBytes)

		fileNameRequest := &Chat{
			Engine: "gpt-3.5-turbo",
			Messages: []Message{
				{
					Role:    "system",
					Content: "You are a helpful assistant.",
				},
				{
					Role:    "user",
					Content: FileNamePrompt + "{\"model\":\"\",\"messages\":[{\"role\":\"system\",\"content\":\"You are a helpful assistant.\"},{\"role\":\"user\",\"content\":\"I love lamp\"},{\"role\":\"assistant\",\"content\":\"That's great to hear! Lamps can add a cozy and decorative touch to any space. Is there anything specific you need help with related to lamps?\"},{\"role\":\"user\",\"content\":\"I love lamp\"},{\"role\":\"assistant\",\"content\":\"I'm glad to hear that you have an appreciation for lamps! Lamps can provide both functionality and aesthetics to a room. If you need any assistance in finding the right type of lamp for your needs or if you have any questions about how to properly maintain and care for your lamps, feel free to ask!\"}]}",
				},
				{
					Role:    "assistant",
					Content: "lamp.yaml",
				},
				{
					Role:    "user",
					Content: FileNamePrompt + jsonString,
				},
			},
		}

		respBody := chat.gptChat(fileNameRequest)
		completionResponse := chat.processResponse(respBody)
		log.Infof("%+v", completionResponse)

		fileName = completionResponse.Choices[0].Message.Content
	} else {
		fileName = strings.Replace(userMessage, "/save ", "", -1)

	}
	err := ToYAML(strings.TrimSpace(fileName), chat)
	if err != nil {
		log.Fatal("Error saving files:", err)
	}
	return "Chat context dumped to file: " + fileName
}
