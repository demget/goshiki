package goshiki

type Forum struct {
	ID        int    `json:"id"`
	Position  int    `json:"position"`
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
	URL       URL    `json:"url"`
}
