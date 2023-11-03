package client

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// Client client.
type Client interface {
	GET(operationID, token string, params interface{}) (map[string]interface{}, error)
	POST(operationID, token string, params interface{}) (map[string]interface{}, error)
}

// httpClient http client.
type httpClient struct {
	url    string
	client *resty.Client
}

// NewClient new client.
func NewClient(url string) Client {
	return &httpClient{
		url:    url,
		client: resty.New(),
	}
}

// GET get unimplemented.
func (c *httpClient) GET(operationID, token string, params interface{}) (map[string]interface{}, error) {
	return nil, nil
}

// POST post.
func (c *httpClient) POST(operationID, token string, params interface{}) (map[string]interface{}, error) {
	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("operationID", operationID).
		SetHeader("token", token).
		SetBody(params).
		Post(c.url)
	if err != nil {
		return nil, err
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(resp.Body(), &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
