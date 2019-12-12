package goshiki

type Video struct {
	ID         int    `json:"id"`
	AnimeID    int    `json:"anime_id"`
	URL        string `json:"url"`
	Source     string `json:"source"`
	Episode    int    `json:"episode"`
	Kind       string `json:"kind"`
	Language   string `json:"language"`
	Quality    string `json:"quality"`
	State      string `json:"state"`
	AuthorName string `json:"author_name"`
}
