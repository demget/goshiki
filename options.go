package goshiki

type Option interface {
	Key() string
}

type (
	Order      string
	Kind       string
	Status     string
	Season     string
	Score      int
	Duration   string
	Rating     string
	Genres     []int
	Studios    []int
	Publishers []int
	Franchises []int
	Censore    bool
	MyList     string
	IDs        []int
	ExcludeIDs []int
	Episode    int
)

const (
	Inbox         string = "inbox"
	Private       string = "private"
	Sent          string = "sent"
	News          string = "news"
	Notifications string = "notifications"

	ByID         Order = "id"
	ByRanked     Order = "ranked"
	ByKind       Order = "kind"
	ByPopularity Order = "popularity"
	ByName       Order = "name"
	ByAiredOn    Order = "aired_on"
	ByEpisodes   Order = "episodes"
	ByVolumes    Order = "volumes"
	ByChapters   Order = "chapters"
	ByStatus     Order = "status"
	ByRandom     Order = "random"

	WithTV      Kind = "tv"
	WithMovie   Kind = "movie"
	WithOVA     Kind = "ova"
	WithONA     Kind = "ona"
	WithSpecial Kind = "special"
	WithMusic   Kind = "music"
	WithTV13    Kind = "tv_13"
	WithTV24    Kind = "tv_24"
	WithTV48    Kind = "tv_48"
	WithManga   Kind = "manga"
	WithManhwa  Kind = "manhwa"
	WithManhua  Kind = "manhua"
	WithNovel   Kind = "novel"
	WithOneshot Kind = "one_shot"
	WithDoujin  Kind = "doujin"
	NotTV       Kind = "!tv"
	NotMovie    Kind = "!movie"
	NotOVA      Kind = "!ova"
	NotONA      Kind = "!ona"
	NotSpecial  Kind = "!special"
	NotMusic    Kind = "!music"
	NotTV13     Kind = "!tv_13"
	NotTV24     Kind = "!tv_24"
	NotTV48     Kind = "!tv_48"
	NotManga    Kind = "!manga"
	NotManhwa   Kind = "!manhwa"
	NotManhua   Kind = "!manhua"
	NotNovel    Kind = "!novel"
	NotOneshot  Kind = "!one_shot"
	NotDoujin   Kind = "!doujin"

	Anons       Status = "anons"
	Ongoing     Status = "ongoing"
	Released    Status = "released"
	NotAnons    Status = "!anons"
	NotOngoing  Status = "!ongoing"
	NotReleased Status = "!released"

	DurationS Duration = "S"
	DurationD Duration = "D"
	DurationF Duration = "F"

	RatingNone  Rating = "none"
	RatingG     Rating = "g"
	RatingPG    Rating = "pg"
	RatingPG13  Rating = "pg_13"
	RatingR     Rating = "r"
	RatingRPlus Rating = "r_plus"
	RatingRX    Rating = "rx"

	Censored Censore = true

	Planned    MyList = "planned"
	Watching   MyList = "watching"
	Rewatching MyList = "rewatching"
	Completed  MyList = "completed"
	OnHold     MyList = "on_hold"
	Dropped    MyList = "dropped"
)

func (Order) Key() string {
	return "order"
}

func (Kind) Key() string {
	return "kind"
}

func (Status) Key() string {
	return "status"
}

func (Season) Key() string {
	return "season"
}

func (Score) Key() string {
	return "score"
}

func (Duration) Key() string {
	return "duration"
}

func (Rating) Key() string {
	return "rating"
}

func (Genres) Key() string {
	return "genre"
}

func (Studios) Key() string {
	return "studio"
}

func (Publishers) Key() string {
	return "publisher"
}

func (Franchises) Key() string {
	return "franchise"
}

func (Censore) Key() string {
	return "censored"
}

func (MyList) Key() string {
	return "mylist"
}

func (IDs) Key() string {
	return "ids"
}

func (ExcludeIDs) Key() string {
	return "exclude_ids"
}

func (Episode) Key() string {
	return "episode"
}
