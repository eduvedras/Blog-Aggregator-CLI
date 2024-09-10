package main

import "fmt"

func commandHelp(_ *config, _ ...string) error {
	commands := getCommands()

	for _, com := range commands {
		fmt.Println("--------------------------------------------------------")
		fmt.Printf("Name:%v\nFormat:%v\nDescription:%v\n", com.name, com.format, com.description)
	}
	return nil
}
