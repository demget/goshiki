package goshiki

import (
	"fmt"
	"net/url"
	"strings"

	urlquery "github.com/google/go-querystring/query"
)

type Request interface {
	Values() url.Values
}

// Params represents an url query as key-value of any type.
// Values of the map converts to string by fmt.Sprint.
type Params map[string]interface{}

// Search represents a request for searching and querying.
// You can also specify page and count of elements.
// If the next page exists, result contains Limit+1 elements.
type Search struct {
	Page  uint   `url:"page,omitempty"`
	Limit uint   `url:"limit,omitempty"`
	Query string `url:"search,omitempty"`
}

// Topics represents a request for searching by topics.
type Topics struct {
	Page    uint   `url:"page,omitempty"`
	Limit   uint   `url:"limit,omitempty"`
	Kind    string `url:"kind,omitempty"`
	Episode int    `url:"episode,omitempty"`
}

// Builder allows you to merge request with options.
type Builder struct {
	Request Request
	Options []Option
}

// Values converts map[string]interface{} to url.Values.
func (ps Params) Values() url.Values {
	values := make(url.Values, len(ps))
	for k, v := range ps {
		values.Set(k, fmt.Sprint(v))
	}
	return values
}

// Values implements Request interface.
func (s Search) Values() url.Values {
	values, _ := urlquery.Values(s)
	return values
}

// Values implements Request interface.
func (b Builder) Values() url.Values {
	values := b.Request.Values()
	for _, opt := range b.Options {
		k, s := opt.Key(), strings.Trim(fmt.Sprint(opt), "[]")
		if v := values.Get(k); v != "" {
			s = v + "," + s
		}
		values.Set(k, s)
	}
	return values
}
