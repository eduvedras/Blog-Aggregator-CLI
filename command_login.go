package main

import (
	"fmt"
	"strings"
)

func commandLogIn(conf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("You provided %v arguments, you should provide only 1 argument which should be the apikey", len(args))
	}

	user, err := conf.blogApiClient.GetUser(args[0])
	if err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return fmt.Errorf("Could not login with apikey %v, user does not exist", args[0])
		}
		return err
	}

	conf.apiKey = user.APIKey
	fmt.Printf("Logged in as %v\n", user.Name)
	return nil
}
