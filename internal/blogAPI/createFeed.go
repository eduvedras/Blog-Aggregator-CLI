package blogapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateFeedResp struct {
	Feed       Feed
	FeedFollow FeedFollow
}

func (c *Client) CreateFeed(key string, name string, feedUrl string) (CreateFeedResp, error) {
	url := baseURL + "/feeds"

	reqBody := struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{
		Name: name,
		Url:  feedUrl,
	}

	out, err := json.Marshal(reqBody)
	if err != nil {
		return CreateFeedResp{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(out))
	if err != nil {
		return CreateFeedResp{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", key))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CreateFeedResp{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateFeedResp{}, err
	}

	feedResp := CreateFeedResp{}
	err = json.Unmarshal(dat, &feedResp)
	if err != nil {
		return CreateFeedResp{}, err
	}

	return feedResp, nil
}
