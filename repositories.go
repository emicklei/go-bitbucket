package bitbucket

import (
	"encoding/json"
	"net/http"
)

func (c *Client) Repository(repo_slug string) (repo Repository, err error) {
	u := c.URLTo("/repositories", c.user, repo_slug)
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return repo, err
	}
	resp, err := c.API.Do(req)
	if err != nil {
		return repo, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&repo)
	return repo, nil
}

func (c *Client) RepositoriesDo(block func(Repository) error) error {
	u := c.URLTo("/repositories", c.user)
	next := u.String()
	for {
		r, err := http.NewRequest("GET", next, nil)
		r.Header.Set("Accept", "application/json")
		if err != nil {
			return err
		}
		resp, err := c.API.Do(r)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var p Page
		json.NewDecoder(resp.Body).Decode(&p)
		for _, each := range p.Repositories() {
			if err := block(each); err != nil {
				return err
			}
		}
		if len(p.Next) == 0 {
			break
		}
		next = c.AuthorizedURL(p.Next).String()
	}
	return nil
}
