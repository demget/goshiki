package goshiki

type Ranobe struct {
	Title
	Score    string `json:"score"`
	Volumes  int    `json:"volumes"`
	Chapters int    `json:"chapters"`
}

type RanobeExt struct {
	Ranobe

	NamesEn            []string          `json:"english"`
	NamesJp            []string          `json:"japanese"`
	Synonyms           []string          `json:"synonyms"`
	LicenseName        string            `json:"license_name_ru"`
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

// Ranobes returns a list of ranobe. You can use Search struct as request.
// Valid options: Order, Kind, Status, Season, Score, Genres,
// Publishers, Franchises, Censored, MyList, IDs, ExcludeIDs.
func (api *API) Ranobes(req Request, opts ...Option) ([]Ranobe, error) {
	values := Builder{req, opts}.Values()

	var ranobes []Ranobe
	if err := api.Raw(EndpointRanobe, values, &ranobes); err != nil {
		return nil, err
	}

	for i := range ranobes {
		ranobes[i].api = api
		ranobes[i].endpoint = EndpointRanobe
	}
	return ranobes, nil
}

// RanobeByID returns an extended ranobe info by its integer ID.
func (api *API) RanobeByID(id int) (RanobeExt, error) {
	var ranobe RanobeExt
	ranobe.api = api
	ranobe.endpoint = EndpointRanobe
	return ranobe, api.Raw(EndpointRanobe.With(id), nil, &ranobe)
}
