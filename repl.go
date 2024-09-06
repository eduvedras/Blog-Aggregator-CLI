package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

func startRepl() {
	commands := getCommands()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Blog-Aggregator > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if command, exists := commands[commandName]; exists {
			err := command.callback(args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(commandName, ": command not found")
		}
	}
}
