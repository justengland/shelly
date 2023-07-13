package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		// If the level is not set or it has a wrong value, default to info level
		level = log.InfoLevel
	}
	log.SetLevel(level)

	fmt.Println("I am shelly your little helper")

	// Add handlers here
	commandHandlers := CommandHandlers{}

	chatHandler := NewChatCommand()
	commandHandlers.SetDefault(chatHandler)
	commandHandlers.AddHandler(NewHelpCommand())
	commandHandlers.AddHandler(NewDumpCommand())
	commandHandlers.AddHandler(chatHandler)

	// Set up a signal channel to capture Ctrl+C (SIGINT)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nReceived Ctrl+C. Exiting...")
		os.Exit(0)
	}()

	// ANSI escape code for yellow
	yellow := "\033[33m"
	// ANSI escape code for green
	green := "\033[32m"
	// ANSI escape code to reset color
	reset := "\033[0m"

	// Chat loop
	for {
		fmt.Print(yellow + "\nUser: " + reset)
		userMessage := getUserInput(commandHandlers)

		assistantMessage := commandHandlers.HandleMessage(userMessage)
		fmt.Println(green + "\nAssistant:" + assistantMessage + reset)
	}
}

func getUserInput(commandHandlers CommandHandlers) string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	// Read loop
	i := 0
	for {
		start := time.Now()
		scanner.Scan()
		next := time.Now()
		// Calculate the time difference in milliseconds
		duration := next.Sub(start)
		ms := duration.Milliseconds()
		line := scanner.Text()

		// If this is a command, just exit dont wait for the return key
		if i == 0 && commandHandlers.IsCommand(line) {
			return strings.TrimSpace(line)
		}

		// Check if this is the first round
		if ms > 1 {
			if strings.ToLower(strings.TrimSpace(line)) == "" {
				break
			}
		}

		lines = append(lines, line)

		i++
	}

	scannerErr := scanner.Err()
	if scannerErr != nil {
		log.Fatal(scannerErr)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}
