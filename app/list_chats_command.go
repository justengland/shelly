package main

func NewListChatsCommand() CommandHandler {
	return CommandHandler{
		Name:    "list_chat",
		Command: listChatsCommand,
	}
}

func listChatsCommand(userMessage string) string {
	return "fail"
}
