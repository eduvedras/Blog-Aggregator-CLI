package blogapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetFeeds(offset int, limit int) ([]Feed, error) {
	url := baseURL + fmt.Sprintf("/feeds?offset=%v&limit=%v", offset, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []Feed{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []Feed{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Feed{}, err
	}

	feeds := []Feed{}
	err = json.Unmarshal(dat, &feeds)
	if err != nil {
		return []Feed{}, err
	}

	return feeds, nil
}
