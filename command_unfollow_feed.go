package main

import (
	"errors"
	"fmt"
)

func commandUnfollowFeed(conf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("You provided %v arguments, you need to provide only one argument which is the feed_followID you want to delete", len(args))
	}

	if conf.apiKey == "" {
		return errors.New("You need to be logged in")
	}

	status, err := conf.blogApiClient.DeleteFeedFollow(conf.apiKey, args[0])
	if err != nil {
		return err
	}

	if status != 200 {
		return errors.New("Invalid Id. You can check the IDs of your feed_follows by using the command `get_feed_follows`.")
	}

	fmt.Printf("feed_follow %v is deleted\n", args[0])
	return nil
}
