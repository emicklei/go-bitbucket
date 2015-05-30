package bitbucket

import (
	"os"
	"testing"
)

func TestRepositories(t *testing.T) {
	c := New(os.Getenv("BB_USER"), os.Getenv("BB_PWD"))
	err := c.RepositoriesDo(func(each Repository) error {
		t.Logf("%s", each.Name)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TestRepository(t *testing.T) {
	c := New(os.Getenv("BB_USER"), os.Getenv("BB_PWD"))
	repo, err := c.Repository("erlang")
	if err != nil {
		t.Error("no such repo")
	}
	t.Logf("%#v", repo)
}
