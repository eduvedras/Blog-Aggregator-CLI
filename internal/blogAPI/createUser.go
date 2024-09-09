package blogapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CreateUser(name string) (User, error) {
	url := baseURL + "/users"
	reqBody := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	out, err := json.Marshal(reqBody)
	if err != nil {
		return User{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(out))
	if err != nil {
		return User{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return User{}, err
	}

	if resp.StatusCode > 299 {
		return User{}, fmt.Errorf("create user command failed with argument %v and status: %v", name, resp.Status)
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
