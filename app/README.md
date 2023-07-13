# Chat Commands

This directory contains the chat commands for the chat assistant.

## Summary

The summary command provides a summary of the chat messages.

## WebSearch

The web search command performs a web search based on the given query.

## Usage

To add new chat commands, create a new Go file in this directory and implement the ChatCommanderFactory interface in the factory.go file. Update the Create method to return the appropriate command handler based on the command.

To register the chat commands dynamically, update the RegisterCommands function in the registry.go file to import the new file and register the factory in the factory map.

