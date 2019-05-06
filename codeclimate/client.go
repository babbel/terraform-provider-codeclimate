package codeclimate

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const codeClimateApiHost string = "https://api.codeclimate.com/v1"

type Client struct {
	apiKey string
}

// TODO: Extend in the future to accept POST requests
func (c *Client) makeRequest(path string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", codeClimateApiHost, path), nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Accept", `W/"application/vnd.api+json"`)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.apiKey))

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}
