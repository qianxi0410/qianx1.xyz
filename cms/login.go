package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

// login is a pre function that is called before the execution func
// we need to get jwt token from the server
// and set it to the request header
func login() (string, error) {
	cli := resty.New()

	resp, err := cli.
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(User{
			Name:     os.Getenv("USER_NAME"),
			Password: os.Getenv("PASSWORD"),
		}).
		Post("https://blogapi.qianx1.xyz/user/login")

	if err != nil || resp.StatusCode() != http.StatusOK {
		return "", err
	}

	var r R
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return "", err
	}

	return r.Data, nil
}
