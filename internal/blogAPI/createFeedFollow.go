package blogapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CreateFeedFollow(key string, feedId string) (FeedFollow, error) {
	url := baseURL + "/feed_follows"

	reqBody := struct {
		FeedID string `json:"feed_id"`
	}{
		FeedID: feedId,
	}

	out, err := json.Marshal(reqBody)
	if err != nil {
		return FeedFollow{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(out))
	if err != nil {
		return FeedFollow{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %v", key))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return FeedFollow{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return FeedFollow{}, err
	}

	feedFollow := FeedFollow{}
	err = json.Unmarshal(dat, &feedFollow)
	if err != nil {
		return FeedFollow{}, err
	}

	return feedFollow, nil
}
