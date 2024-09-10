package main

import (
	"errors"
	"fmt"
)

func commandGetPosts(conf *config, args ...string) error {
	n := len(args)
	var limit int
	var err error
	switch n {
	case 0:
		limit = 10
	case 2:
		if args[0] == "--limit" {
			limit, err = checkOptionalParam(0, args...)
			if err != nil {
				return err
			}
			break
		}
		return fmt.Errorf("Your provided flag is '%v', you need the flag '--limit'", args[0])
	default:
		return errors.New("You provided the wrong amount of arguments, use `help` to see this command format")
	}

	if conf.apiKey == "" {
		return errors.New("You need to be logged in")
	}

	posts, err := conf.blogApiClient.GetPosts(conf.apiKey, limit)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Printf("Id:%v\nFeedID:%v\nTitle:%v\nDescription:%v\nUrl:%v\nPublishedAt:%v\n", post.ID, post.FeedID, post.Title, post.Description, post.URL, post.PublishedAt)
	}

	return nil
}
