package goshiki

import (
	"encoding/json"
	"time"
)

type Topic struct {
	ID                int       `json:"id"`
	Title             string    `json:"topic_title"`
	Body              string    `json:"body"`
	BodyHTML          string    `json:"html_body"`
	FooterHTML        string    `json:"html_footer"`
	CreatedAt         time.Time `json:"created_at"`
	Comments          int       `json:"comments_count"`
	Forum             Forum     `json:"forum"`
	User              User      `json:"user"`
	Type              string    `json:"type"`
	Linked            Linked    `json:"-"`
	IsViewed          bool      `json:"viewed"`
	LastViewedComment *Comment  `json:"last_comment_viewed"`
	Event             string    `json:"event"`

	//? Episode interface{} `json:"episode"`
}

func (t *Topic) UnmarshalJSON(data []byte) error {
	type topic Topic

	var aux topic
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	*t = Topic(aux)
	return t.Linked.UnmarshalJSON(data)
}
