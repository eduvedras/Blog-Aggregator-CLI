package main

import "fmt"

func commandHelp(_ *config, _ ...string) error {
	fmt.Println("Help message")
	return nil
}
