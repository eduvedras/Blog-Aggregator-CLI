package main

import (
	"errors"
	"fmt"
)

func commandFollowFeed(conf *config, args ...string) error {
	if conf.apiKey == "" {
		return errors.New("You need to be logged in")
	}

	if len(args) != 1 {
		return fmt.Errorf("You provided %v arguments, you need to provide only one which is the feedID", len(args))
	}

	feedFollow, err := conf.blogApiClient.CreateFeedFollow(conf.apiKey, args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Followed feed with ID \"%v\".", feedFollow.FeedID)
	return nil
}
