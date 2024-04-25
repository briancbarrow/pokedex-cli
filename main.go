package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := initializeCommands()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(printPrompt())
	for scanner.Scan() {
		text := scanner.Text()
		commands[text].callback()
		fmt.Print(printPrompt())
	}
}

func initializeCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func printPrompt() string {
	return "Pokedex CLI > "
}

func commandHelp() error {
	commands := initializeCommands()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s", cmd.name, cmd.description)
		fmt.Println()
	}
	fmt.Println("")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
