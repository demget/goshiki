package goshiki

import (
	"time"

	urlquery "github.com/google/go-querystring/query"
)

type Title struct {
	ID         int    `json:"id"`
	NameEn     string `json:"name"`
	NameRu     string `json:"russian"`
	Image      Image  `json:"image"`
	URL        URL    `json:"url"`
	Kind       string `json:"kind"`
	Status     string `json:"status"`
	AiredOn    Time   `json:"aired_on"`
	ReleasedOn Time   `json:"released_on"`

	api      *API
	endpoint Endpoint
}

func (title Title) getAPI() *API {
	if title.api == nil {
		panic("goshiki: API is nil")
	}
	return title.api
}

type Role struct {
	RolesEn   []string `json:"roles"`
	RolesRu   []string `json:"roles_russian"`
	Character *Base    `json:"character"`
	Person    *Base    `json:"person"`
}

type Franchise struct {
	Links     []FranchiseLink `json:"links"`
	Nodes     []FranchiseNode `json:"nodes"`
	CurrentID int             `json:"current_id"`
}

type FranchiseLink struct {
	ID       int    `json:"id"`
	SourceID int    `json:"source_id"`
	TargetID int    `json:"target_id"`
	Source   int    `json:"source"`
	Target   int    `json:"target"`
	Weight   int    `json:"weight"`
	Relation string `json:"relation"`
}

type FranchiseNode struct {
	ID       int    `json:"id"`
	Date     int    `json:"date"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	URL      string `json:"url"`
	Year     int    `json:"year"`
	Kind     string `json:"kind"`
	Weight   int    `json:"weight"`
}

type ExternalLink struct {
	ID         int        `json:"id"`
	Kind       string     `json:"kind"`
	URL        string     `json:"url"`
	Source     string     `json:"source"`
	EntryID    int        `json:"entry_id"`
	EntryType  string     `json:"entry_type"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	ImportedAt *time.Time `json:"imported_at"`
}

// Roles returns title's roles.
func (title Title) Roles() ([]Role, error) {
	var roles []Role
	endpoint := title.endpoint.With(title.ID, "roles")
	return roles, title.getAPI().Raw(endpoint, nil, &roles)
}

// Franchise returns title's franchise.
func (title Title) Franchise() (Franchise, error) {
	var franchise Franchise
	endpoint := title.endpoint.With(title.ID, "franchise")
	return franchise, title.getAPI().Raw(endpoint, nil, &franchise)
}

// ExternalLinks returns title's external links.
func (title Title) ExternalLinks() ([]ExternalLink, error) {
	var links []ExternalLink
	endpoint := title.endpoint.With(title.ID, "external_links")
	return links, title.getAPI().Raw(endpoint, nil, &links)
}

// Topics returns title's topics.
func (title Title) Topics(req Topics) ([]Topic, error) {
	values, _ := urlquery.Values(req)

	var topics []Topic
	endpoint := title.endpoint.With(title.ID, "topics")
	if err := title.getAPI().Raw(endpoint, values, &topics); err != nil {
		return nil, err
	}

	for i := range topics {
		topics[i].User.api = title.api
	}
	return topics, nil
}
