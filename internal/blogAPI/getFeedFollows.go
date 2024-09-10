package blogapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetFeedFollows(key string) ([]FeedFollow, error) {
	url := baseURL + "/feed_follows"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []FeedFollow{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %v", key))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []FeedFollow{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return []FeedFollow{}, err
	}

	feedFollows := []FeedFollow{}
	err = json.Unmarshal(dat, &feedFollows)
	if err != nil {
		return []FeedFollow{}, err
	}

	return feedFollows, nil
}
