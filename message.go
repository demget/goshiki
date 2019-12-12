package goshiki

import (
	"encoding/json"
	"time"
)

type Message struct {
	ID        int       `json:"id"`
	Kind      string    `json:"kind"`
	IsRead    bool      `json:"read"`
	Body      string    `json:"body"`
	BodyHTML  string    `json:"html_body"`
	CreatedAt time.Time `json:"created_at"`
	Linked    Linked    `json:"-"`
	From      User      `json:"from"`
	To        User      `json:"to"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type message Message

	var aux message
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	*m = Message(aux)
	return m.Linked.UnmarshalJSON(data)
}
