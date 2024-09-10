# Blog-Aggregator-CLI

Blog-Aggregator-CLI is command-line interface that helps abstract the interaction with the [Blog-Aggregator API](https://github.com/eduvedras/Blog-Aggregator). An API that allows users to aggregate all their favorite RSS blogs and feeds in one place.

## Quickstart
Run the following commands on your terminal to start.

```bash
git clone https://github.com/eduvedras/Blog-Aggregator-CLI.git
cd Blog-Aggregator-CLI
make brun
./blog-aggregator-cli
```

## Usage
The following is a brief overview of the commands.

```bash
--------------------------------------------------------
Name:get_posts
Format:get_posts [--limit <number>]
Description:Get the posts of the feeds followed by the user. There is an optional limit parameter to specify the number of posts you want to list, it defaults to 10.
--------------------------------------------------------
Name:new_user
Format:new_user <name>
Description:Creates a new user
--------------------------------------------------------
Name:login
Format:login <apikey>
Description:Log in into a user account
--------------------------------------------------------
Name:get_feeds
Format:get_feeds [--offset <number>] [--limit <number>]
Description:Get all feeds. You can provide two optional parameters offset and limit, the offset the position where you want to start to list the feeds and limit is the number of feeds you want to list. Offset defaults to 0 and limit defaults to 20.
--------------------------------------------------------
Name:unfollow_feed
Format:unfollow_feed <feed_follow_ID>
Description:Unfollow a feed
--------------------------------------------------------
Name:get_feed_follows
Format:get_feed_follows
Description:Get the feed_follows of the user
--------------------------------------------------------
Name:help
Format:help
Description:Displays a help message
--------------------------------------------------------
Name:healthz
Format:healthz
Description:Check the status of the server
--------------------------------------------------------
Name:new_feed
Format:new_feed <name> <url>
Description:Create a new feed
--------------------------------------------------------
Name:follow_feed
Format:follow_feed <feedID>
Description:Follow an existing feed
```
