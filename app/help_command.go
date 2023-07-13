package main

import (
	"fmt"
)

func NewHelpCommand() CommandHandler {
	return CommandHandler{
		Name:    "help",
		Command: helpCommand,
	}
}

func helpCommand(userMessage string) string {
	return fmt.Sprintf(`
	Welcome to shelly terminal application!

	Here are some commands you can use:

	- chat: Enter the chat interface.
	- summarize: Summarize the recent chats.
	- help: Display this help message.
	- exit: Exit the application.

	Remember, you can always type 'help' to get help on using this application. Enjoy!
	`)
}
