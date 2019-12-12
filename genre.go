package goshiki

type Genre struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	NameRu string `json:"russian"`
	Kind   string `json:"kind"`
}
