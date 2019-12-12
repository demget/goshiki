package goshiki

type Manga struct {
	Title
	Volumes  int `json:"volumes"`
	Chapters int `json:"chapters"`
}

type MangaExt struct {
	Manga

	NamesEn            []string          `json:"english"`
	NamesJp            []string          `json:"japanese"`
	Synonyms           []string          `json:"synonyms"`
	LicenseName        string            `json:"license_name_ru"`
	Score              string            `json:"score"`
	Description        string            `json:"description"`
	DescriptionHTML    string            `json:"description_html"`
	DescriptionSource  string            `json:"description_source"`
	FranchiseKey       string            `json:"franchise"`
	IsFavoured         bool              `json:"favoured"`
	IsAnons            bool              `json:"anons"`
	IsOngoing          bool              `json:"ongoing"`
	ThreadID           int               `json:"thread_id"`
	TopicID            int               `json:"topic_id"`
	MyanimelistID      int               `json:"myanimelist_id"`
	RatesScoresStats   []IntNameValue    `json:"rates_scores_stats"`
	RatesStatusesStats []StringNameValue `json:"rates_statuses_stats"`
	Genres             []Genre           `json:"genres"`
	Publishers         []Publisher       `json:"publishers"`
	UserRate           Rate              `json:"user_rate"`
}

// Mangas returns a list of mangas. You can use Search struct as request.
// Valid options: Order, Kind, Status, Season, Score, Genres,
// Publishers, Franchises, Censored, MyList, IDs, ExcludeIDs.
func (api *API) Mangas(req Request, opts ...Option) ([]Manga, error) {
	values := Builder{req, opts}.Values()

	var mangas []Manga
	if err := api.Raw(EndpointMangas, values, &mangas); err != nil {
		return nil, err
	}

	for i := range mangas {
		mangas[i].api = api
		mangas[i].endpoint = EndpointMangas
	}
	return mangas, nil
}

// MangaByID returns an extended manga info by its integer ID.
func (api *API) MangaByID(id int) (MangaExt, error) {
	var manga MangaExt
	manga.api = api
	manga.endpoint = EndpointMangas
	return manga, api.Raw(EndpointMangas.With(id), nil, &manga)
}
