package main

import "fmt"

func commandHelp(_ ...string) error {
	fmt.Println("Help message")
	return nil
}
