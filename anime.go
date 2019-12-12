package goshiki

import "time"

type Anime struct {
	Title
	Episodes      int `json:"episodes"`
	AiredEpisodes int `json:"episodes_aired"`
}

type AnimeExt struct {
	Anime

	Rating             Rating            `json:"rating"`
	NamesEn            []string          `json:"english"`
	NamesJp            []string          `json:"japanese"`
	Synonyms           []string          `json:"synonyms"`
	LicenseName        string            `json:"license_name_ru"`
	Duration           int               `json:"duration"`
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
	UpdatedAt          time.Time         `json:"updated_at"`
	NextEpisodeAt      time.Time         `json:"next_episode_at"`
	Genres             []Genre           `json:"genres"`
	Studios            []Studio          `json:"studios"`
	Videos             []AnimeVideo      `json:"videos"`
	UserRate           Rate              `json:"user_rate"`
	Screenshots        []Image           `json:"screenshots"`
}

type AnimeVideo struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	ImageURL  string `json:"image_url"`
	PlayerURL string `json:"player_url"`
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Hosting   string `json:"hosting"`
}

type AnimeRole struct {
	RolesEn   []string `json:"roles"`
	RolesRu   []string `json:"roles_russian"`
	Character *Base    `json:"character"`
	Person    *Base    `json:"person"`
}

type AnimeTopic struct {
	Anime

	ThreadID int    `json:"thread_id"`
	TopicID  int    `json:"topic_id"`
	TopicURL string `json:"topic_url"`
	Type     string `json:"type"`
}

// Animes returns a list of animes. You can use Search struct as request.
// Valid options: Order, Kind, Status, Season, Score, Duration, Rating,
// Genres, Studios, Franchises, Censore, MyList, IDs, ExcludeIDs.
func (api *API) Animes(req Request, opts ...Option) ([]Anime, error) {
	values := Builder{req, opts}.Values()

	var animes []Anime
	if err := api.Raw(EndpointAnimes, values, &animes); err != nil {
		return nil, err
	}

	for i := range animes {
		animes[i].api = api
		animes[i].endpoint = EndpointAnimes
	}
	return animes, nil
}

// AnimeByID returns an extended anime info by its integer ID.
func (api *API) AnimeByID(id int) (AnimeExt, error) {
	var anime AnimeExt
	anime.api = api
	anime.endpoint = EndpointAnimes
	return anime, api.Raw(EndpointAnimes.With(id), nil, &anime)
}

// Similar returns similar animes.
func (anime Anime) Similar() ([]Anime, error) {
	var animes []Anime
	endpoint := EndpointAnimes.With(anime.ID, "similar")
	if err := anime.getAPI().Raw(endpoint, nil, &animes); err != nil {
		return nil, err
	}

	for i := range animes {
		animes[i].api = anime.api
		animes[i].endpoint = EndpointAnimes
	}
	return animes, nil
}

// Related returns related animes.
func (anime Anime) Related() ([]Anime, error) {
	var animes []Anime
	endpoint := EndpointAnimes.With(anime.ID, "related")
	if err := anime.getAPI().Raw(endpoint, nil, &animes); err != nil {
		return nil, err
	}

	for i := range animes {
		animes[i].api = anime.api
		animes[i].endpoint = EndpointAnimes
	}
	return animes, nil
}

// AllScreenshots returns a list of all anime screenshots.
func (anime Anime) AllScreenshots() ([]Image, error) {
	var screens []Image
	endpoint := EndpointAnimes.With(anime.ID, "screenshots")
	return screens, anime.getAPI().Raw(endpoint, nil, &screens)
}
