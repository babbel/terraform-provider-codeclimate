package codeclimateclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const codeClimateApiHost string = "https://api.codeclimate.com/v1"

type Client struct {
	ApiKey  string
	BaseUrl string
}

// TODO: Extend in the future to accept POST requests
func (c *Client) makeRequest(path string) ([]byte, error) {
	client := &http.Client{}

	if c.BaseUrl == "" {
		c.BaseUrl = codeClimateApiHost
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.BaseUrl, path), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.ApiKey))

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}
