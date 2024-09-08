package blogapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetUser(key string) (User, error) {
	url := baseURL + "/users"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", key))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return User{}, err
	}

	if resp.StatusCode > 299 {
		return User{}, fmt.Errorf("Could not get user with apikey %v, request failed with status %v", key, resp.Status)
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	user := User{}
	err = json.Unmarshal(dat, &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
