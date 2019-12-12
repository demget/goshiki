package goshiki

type Character struct {
	Base
	Altname           string `json:"altname"`
	NameJp            string `json:"japanese"`
	Description       string `json:"description"`
	DescriptionHTML   string `json:"description_html"`
	DescriptionSource string `json:"description_source"`
	IsFavoured        bool   `json:"favoured"`
	ThreadID          int    `json:"thread_id"`
	TopicID           int    `json:"topic_id"`
	UpdatedAt         Time   `json:"updated_at"`
	Seyu              []Base `json:"seyu"`

	Animes []struct {
		Anime
		Roles []string `json:"roles"`
		Role  string   `json:"role"`
	} `json:"animes"`
	Mangas []struct {
		Manga
		Roles []string `json:"roles"`
		Role  string   `json:"role"`
	} `json:"mangas"`
}
