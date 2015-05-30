package bitbucket

import "github.com/mitchellh/mapstructure"

type Page struct {
	Size   int64  `json:"size"`    // Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Number int    `json:"page"`    // Page number of the current results. This is an optional element that is not provided in all responses.
	Length int    `json:"pagelen"` // Current number of objects on the existing page. Globally, the minimum length is 10 and the maximum is 100. Some APIs may specify a different default.
	Next   string `json:"next"`    // Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	// Link toe previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available.Client
	// Use this link to navigate the result set and refrain from constructing your own URLs.
	Previous string        `json:"previous"`
	Values   []interface{} `json:"values"` // The list of objects. This contains at most pagelen objects.
}

func (p Page) Repositories() []Repository {
	list := []Repository{}
	for _, each := range p.Values {
		var r Repository
		mapstructure.Decode(each, &r)
		list = append(list, r)
	}
	return list
}
