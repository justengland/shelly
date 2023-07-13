package main

var chat = &Chat{
	Messages: []Message{
		Message{
			Role:    "system",
			Content: "You are a helpful assistant.",
		},
	},
}

func NewChatCommand() CommandHandler {
	return CommandHandler{
		Name:    "chat",
		Command: chatCommand,
	}
}

func chatCommand(userMessage string) string {

	chat.Messages = append(chat.Messages, Message{Role: "user", Content: userMessage})

	completionRequest := &Chat{
		Engine:   "gpt-3.5-turbo",
		Messages: chat.Messages,
	}

	respBody := chat.gptChat(completionRequest)
	completionResponse := chat.processResponse(respBody)

	chat.Messages = append(chat.Messages, completionResponse.Choices[0].Message)

	return completionResponse.Choices[0].Message.Content
}
