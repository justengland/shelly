package main

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type CommandHandler struct {
	Name    string
	Command func(userMessage string) string
}

type CommandHandlers struct {
	Handlers       []CommandHandler
	DefaultHandler CommandHandler
}

func (c *CommandHandlers) SetDefault(handler CommandHandler) {
	c.DefaultHandler = handler
}

func (c *CommandHandlers) AddHandler(handler CommandHandler) {
	c.Handlers = append(c.Handlers, handler)
}

func (c *CommandHandlers) HandleMessage(message string) string {
	for _, handler := range c.Handlers {
		if getHandler(handler.Name, message) {
			log.Infof("handler found: " + handler.Name)
			return handler.Command(message)
		}
	}

	log.Infof("default handler: " + c.DefaultHandler.Name)

	if c.DefaultHandler.Command != nil {
		return c.DefaultHandler.Command(message)
	}

	return "todo make a default chat handler"
}

func (c *CommandHandlers) IsCommand(message string) bool {
	for _, handler := range c.Handlers {
		if getHandler(handler.Name, message) {
			log.Infof("handler found: " + handler.Name)
			return true
		}
	}

	return false
}

func getHandler(name string, message string) bool {
	return strings.HasPrefix(message, "/"+name)
}
