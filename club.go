package goshiki

type Club struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Logo          ClubImage `json:"logo"`
	IsCensored    bool      `json:"is_censored"`
	JoinPolicy    string    `json:"join_policy"`
	CommentPolicy string    `json:"comment_policy"`
}

type ClubImage struct {
	Original string `json:"original"`
	Main     string `json:"main"`
	X96      string `json:"x96"`
	X73      string `json:"x73"`
	X48      string `json:"x48"`
}
