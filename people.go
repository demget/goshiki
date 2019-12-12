package goshiki

import "time"

type Person struct {
	Base
	NameJp             string    `json:"japanese"`
	JobTitle           string    `json:"job_title"`
	Birthday           Time      `json:"birthday"`
	Website            string    `json:"website"`
	ThreadID           int       `json:"thread_id"`
	TopicID            int       `json:"topic_id"`
	IsPersonFavoured   bool      `json:"person_favoured"`
	IsProducer         bool      `json:"producer"`
	IsProducerFavoured bool      `json:"producer_favoured"`
	IsMangaka          bool      `json:"mangaka"`
	IsMangakaFavoured  bool      `json:"mangaka_favoured"`
	IsSeyu             bool      `json:"seyu"`
	IsSeyuFavoured     bool      `json:"seyu_favoured"`
	UpdatedAt          time.Time `json:"updated_at"`

	//? GrouppedRoles [][]interface{} `json:"groupped_roles"`

	Roles []struct {
		Characters []Base  `json:"characters"`
		Animes     []Anime `json:"animes"`
	} `json:"roles"`
	Works []struct {
		Anime *Anime `json:"anime"`
		Manga *Manga `json:"manga"`
		Role  string `json:"role"`
	} `json:"works"`
}
