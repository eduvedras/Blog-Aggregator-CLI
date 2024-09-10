package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	format      string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			format:      "help",
			callback:    commandHelp,
		},
		"healthz": {
			name:        "healthz",
			description: "Check the status of the server",
			format:      "healthz",
			callback:    commandHealthz,
		},
		"new_user": {
			name:        "new_user",
			description: "Creates a new user",
			format:      "new_user <name>",
			callback:    commandCreateUser,
		},
		"login": {
			name:        "login",
			description: "Log in into a user account",
			format:      "login <apikey>",
			callback:    commandLogIn,
		},
		"new_feed": {
			name:        "new_feed",
			description: "Create a new feed",
			format:      "new_feed <name> <url>",
			callback:    commandCreateFeed,
		},
		"get_feeds": {
			name:        "get_feeds",
			description: "Get all feeds. You can provide two optional parameters offset and limit, the offset the position where you want to start to list the feeds and limit is the number of feeds you want to list. Offset defaults to 0 and limit defaults to 20.",
			format:      "get_feeds [--offset <number>] [--limit <number>]",
			callback:    commandGetFeeds,
		},
		"follow_feed": {
			name:        "follow_feed",
			description: "Follow an existing feed",
			format:      "follow_feed <feedID>",
			callback:    commandFollowFeed,
		},
		"unfollow_feed": {
			name:        "unfollow_feed",
			description: "Unfollow a feed",
			format:      "unfollow_feed <feed_follow_ID>",
			callback:    commandUnfollowFeed,
		},
		"get_feed_follows": {
			name:        "get_feed_follows",
			description: "Get the feed_follows of the user",
			format:      "get_feed_follows",
			callback:    commandGetFeedFollows,
		},
		"get_posts": {
			name:        "get_posts",
			description: "Get the posts of the feeds followed by the user. There is an optional limit parameter to specify the number of posts you want to list, it defaults to 10.",
			format:      "get_posts [--limit <number>]",
			callback:    commandGetPosts,
		},
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

func startRepl(conf *config) {
	commands := getCommands()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Blog-Aggregator > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if command, exists := commands[commandName]; exists {
			err := command.callback(conf, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(commandName, ": command not found")
		}
	}
}
