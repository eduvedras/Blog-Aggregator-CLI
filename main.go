package main

import (
	"time"

	blogapi "github.com/eduvedras/Blog-Aggregator-CLI/internal/blogAPI"
)

type config struct {
	blogApiClient blogapi.Client
	apiKey        string
}

func main() {
	conf := &config{
		blogApiClient: blogapi.NewClient(5 * time.Second),
	}
	startRepl(conf)
}
