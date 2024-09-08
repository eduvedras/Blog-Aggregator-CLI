package blogapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type status struct {
	Status string `json:"status"`
}

func (c *Client) CheckHealth() (status, error) {
	url := baseURL + "/healthz"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return status{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return status{}, err
	}

	if resp.StatusCode > 299 {
		return status{}, fmt.Errorf("health command failed with status %v:", resp.Status)
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return status{}, err
	}

	respHealth := status{}
	err = json.Unmarshal(dat, &respHealth)
	if err != nil {
		return status{}, err
	}

	return respHealth, err
}
