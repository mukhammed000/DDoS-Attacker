package client

import (
    "net/http"
)

type Client struct {
    http.Client
}

func NewClient() *Client {
    return &Client{
        Client: http.Client{
            Timeout: 10,
        },
    }
}

func (c *Client) SendRequest(url string) (*http.Response, error) {
    resp, err := c.Get(url)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
