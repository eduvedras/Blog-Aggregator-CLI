package main

import "fmt"

func commandHealthz(conf *config, _ ...string) error {
	statusResp, err := conf.blogApiClient.CheckHealth()
	if err != nil {
		return err
	}
	fmt.Println(statusResp.Status)
	return nil
}
