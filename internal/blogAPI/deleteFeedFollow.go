package blogapi

import (
	"fmt"
	"net/http"
)

func (c *Client) DeleteFeedFollow(key string, id string) (int, error) {
	url := baseURL + fmt.Sprintf("/feed_follows/%v", id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %v", key))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
