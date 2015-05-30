package bitbucket

import "time"

type Repository struct {
	SCM         string        `json:"scm"`         //	The source control manager for the repository. This is either hg or git.
	HasWiki     bool          `json:"has_wiki"`    //	A boolean indicating if the repository has a wiki.
	Description string        `json:"description"` //	A string containing the repository's description.
	Links       []interface{} `json:"links"`       //	An array of related objects.
	UpdatedOn   time.Time     `json:"updated_on"`  //	A date timestamp of the last update to this repository.
	ForkPolicy  string        `json:"fork_policy"`
	/*
	   Control the rules for forking this repository. Available values are:

	   allow_forks: unrestricted forking
	   no_public_forks: restrict forking to private forks (forks cannot be made public later)
	   no_forks: deny all forking
	*/
	CreatedOn time.Time `json:"created_on"` //	An ISO-8601 date timestamp of this repository's creation date.
	Owner     string    `json:"owner"`      //	The owner's account.
	Size      int64     `json:"size"`       //    The size of the repository in bytes.
	Parent    string    `json:"parent"`     //	The parent repository this repository was forked off (only present on forks). This is a repository object itself.
	HasIssues bool      `json:"has_issues"` //	A boolean indicating a repository has an issue tracker.
	IsPrivate bool      `json:"is_private"` //	A boolean indicating if a repository is private or public.
	Fullname  string    `json:"full_name"`  //	The unique key into the repository. This key has the format: {owner}/{repo_slug}
	Name      string    `json:"name"`       //	The display name of the repository.
	Language  string    `json:"language"`   //	The main (programming) language of the repository source files.
}
