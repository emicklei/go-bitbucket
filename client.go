package bitbucket

import (
	"net/http"
	"net/url"
	"path/filepath"
)

type Client struct {
	user, password string
	API            *http.Client
}

func New(user, password string) *Client {
	return &Client{
		user:     user,
		password: password,
		API:      new(http.Client),
	}
}

func (c Client) URLTo(pathsegments ...string) *url.URL {
	u := new(url.URL)
	u.Host = "api.bitbucket.org"
	u.Scheme = "https"
	u.User = url.UserPassword(c.user, c.password)
	u.Path = filepath.Join(append([]string{"2.0"}, pathsegments...)...)
	return u
}

func (c Client) AuthorizedURL(unauthorizedURL string) *url.URL {
	u, _ := url.Parse(unauthorizedURL)
	u.User = url.UserPassword(c.user, c.password)
	return u
}
