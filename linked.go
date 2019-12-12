package goshiki

import (
	"encoding/json"
)

type Linked struct {
	ID   int
	Type string

	Anime  *Anime
	Manga  *Manga
	Ranobe *Ranobe
	Topic  *AnimeTopic
}

func (l Linked) setAPI(api *API) {
	switch l.Type {
	case "Review":
		fallthrough
	case "Anime":
		l.Anime.api = api
	case "Manga":
		l.Manga.api = api
	case "Ranobe":
		l.Ranobe.api = api
	case "Topic":
		l.Topic.api = api
	}
}

func (l *Linked) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID   int             `json:"linked_id"`
		Type string          `json:"linked_type"`
		Raw  json.RawMessage `json:"linked"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	*l = Linked{
		ID:   aux.ID,
		Type: aux.Type,
	}

	switch l.Type {
	case "Review":
		fallthrough
	case "Anime":
		return json.Unmarshal(aux.Raw, &l.Anime)
	case "Manga":
		return json.Unmarshal(aux.Raw, &l.Manga)
	case "Ranobe":
		return json.Unmarshal(aux.Raw, &l.Ranobe)
	case "Topic":
		return json.Unmarshal(aux.Raw, &l.Topic)
	}

	return nil
}
