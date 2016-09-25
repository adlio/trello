package trello

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const DEFAULT_BASEURL = "https://api.trello.com/1"

type Client struct {
	client  *http.Client
	BaseURL string
	Key     string
	Token   string
}

func NewClient(key, token string) *Client {
	return &Client{
		client:  http.DefaultClient,
		BaseURL: DEFAULT_BASEURL,
		Key:     key,
		Token:   token,
	}
}

func (c *Client) Get(path string, args Arguments, target interface{}) error {

	params := args.ToURLValues()

	if c.Key != "" {
		params.Set("key", c.Key)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		return errors.Wrapf(err, "Invalid GET request %s", url)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(target)
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed on %s", url)
	}

	return nil
}
