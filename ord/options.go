package ord

import "net/http"

type Option func(a *Client) error

func WithToken(token string) Option {
	return func(c *Client) error {
		c.token = token

		return nil
	}
}

func WithHttpClient(cl *http.Client) Option {
	return func(c *Client) error {
		c.http = cl

		return nil
	}
}

func WithBase(base string) Option {
	return func(c *Client) error {
		c.base = base

		return nil
	}
}
