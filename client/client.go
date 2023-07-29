package client

import (
	"encoding/json"
	"github.com/ifubar/wow/server"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (c Client) GetTask() (server.GetTaskResponse, error) {
	tResp := server.GetTaskResponse{}

	resp, err := http.Get(c.url + "/task")
	if err != nil {
		return tResp, err
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return tResp, err
	}
	err = resp.Body.Close()
	if err != nil {
		return tResp, err
	}

	return tResp, json.Unmarshal(raw, &tResp)
}

func (c Client) GetWisdom(solution, token string) (string, error) {
	req, err := http.NewRequest("POST", c.url+"/wisdom", strings.NewReader(solution))
	if err != nil {
		return "", err
	}
	req.Header[server.JwtHeader] = []string{token}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(raw), nil
}
