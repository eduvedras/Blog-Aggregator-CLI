package main

import (
	"errors"
	"fmt"
)

func commandGetFeedFollows(conf *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("You provided %v arguments, this command takes no arguments", len(args))
	}

	if conf.apiKey == "" {
		return errors.New("You need to be logged in")
	}

	feedFollows, err := conf.blogApiClient.GetFeedFollows(conf.apiKey)
	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollows {
		fmt.Println("---------------------------------------------------------------------------------")
		fmt.Printf("ID:%v\nFeedID:%v\nFollowedAt:%v\n", feedFollow.ID, feedFollow.FeedID, feedFollow.CreatedAt)
	}
	return nil
}
