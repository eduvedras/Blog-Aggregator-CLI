package main

import (
	"fmt"
	"strconv"
)

func checkOptionalParam(pos int, args ...string) (int, error) {
	param, err := strconv.Atoi(args[pos+1])
	if err != nil {
		return 0, fmt.Errorf("Argument provided for '%v' needs to be an integer.", args[pos])
	}
	return param, nil
}

func queryParams(args ...string) (int, int, error) {
	var offset int
	var limit int
	var err error
	flagLimit := "--limit"
	flagOffset := "--offset"

	n := len(args)
	switch n {
	case 0:
		offset = 0
		limit = 20
	case 2:
		if args[0] == flagLimit {
			limit, err = checkOptionalParam(0, args...)
			if err != nil {
				return 0, 0, err
			}
			offset = 0
			break
		}
		if args[0] == flagOffset {
			offset, err = checkOptionalParam(0, args...)
			if err != nil {
				return 0, 0, err
			}
			limit = 20
			break
		}
		return 0, 0, fmt.Errorf("Your provided flag is '%v', you need to provide one of these two '--limit' or '--offset'", args[0])
	case 4:
		if args[0] == flagLimit && args[2] == flagOffset {
			limit, err = checkOptionalParam(0, args...)
			if err != nil {
				return 0, 0, err
			}
			offset, err = checkOptionalParam(2, args...)
			if err != nil {
				return 0, 0, err
			}
			break
		}
		if args[0] == flagOffset && args[2] == flagLimit {
			offset, err = checkOptionalParam(0, args...)
			if err != nil {
				return 0, 0, err
			}
			limit, err = checkOptionalParam(2, args...)
			if err != nil {
				return 0, 0, err
			}
			break
		}
		return 0, 0, fmt.Errorf("Your provided flags are '%v' and '%v', the correct format is '--offset' and '--limit'", args[0], args[2])
	default:
		return 0, 0, fmt.Errorf("You provided %v arguments, to use the optional parameters you need to use the flag followed by only one argument.", n)
	}
	return offset, limit, nil
}

func commandGetFeeds(conf *config, args ...string) error {
	offset, limit, err := queryParams(args...)
	if err != nil {
		return err
	}

	feeds, err := conf.blogApiClient.GetFeeds(offset, limit)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("FeedID:\"%v\", Name: \"%v\", Url:\"%v\", AuthorId:\"%v\"\n", feed.ID, feed.Name, feed.URL, feed.UserID)
	}
	return nil
}
