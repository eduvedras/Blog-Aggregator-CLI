package main

import (
	"errors"
	"fmt"
)

func commandCreateUser(conf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("To create a user you need to provide a name")
	}

	name := ""
	for i, el := range args {
		if i == 0 {
			name = name + el
			continue
		}
		name = name + " " + el
	}

	respUser, err := conf.blogApiClient.CreateUser(name)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Created user with name %v and apikey %v", respUser.Name, respUser.APIKey))

	return nil
}
