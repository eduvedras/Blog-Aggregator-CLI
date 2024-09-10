package blogapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPosts(key string, limit int) ([]Post, error) {
	url := baseURL + fmt.Sprintf("/posts?limit=%v", limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []Post{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %v", key))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []Post{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Post{}, err
	}

	posts := []Post{}
	err = json.Unmarshal(dat, &posts)
	if err != nil {
		return []Post{}, err
	}

	return posts, nil
}
