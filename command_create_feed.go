package main

import (
	"errors"
	"fmt"
)

func commandCreateFeed(conf *config, args ...string) error {
	if conf.apiKey == "" {
		return errors.New("You need to be logged in")
	}

	if len(args) != 2 {
		return fmt.Errorf("You provided %v arguments. You need to provide 2 arguments the name and the url of the feed.", len(args))
	}

	resp, err := conf.blogApiClient.CreateFeed(conf.apiKey, args[0], args[1])
	if err != nil {
		return fmt.Errorf("Failed to created feed with error: %v", err)
	}

	fmt.Printf("Created and followed feed with name %v and url %v.\n", resp.Feed.Name, resp.Feed.URL)
	return nil
}
