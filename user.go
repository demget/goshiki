package goshiki

import "time"

type User struct {
	ID           int       `json:"id"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Image        Image     `json:"image"`
	LastOnlineAt time.Time `json:"last_online_at"`

	api *API
}

func (user User) getAPI() *API {
	if user.api == nil {
		panic("goshiki: API is nil")
	}
	return user.api
}

type UserInfo struct {
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Website string `json:"website"`
	BirthOn *Time  `json:"birth_on"`
	Locale  string `json:"locale"`
}

type UserExt struct {
	// BirthOn and Locale will always be default (nil and empty string).
	// For this kind of data use user.Info().
	UserInfo

	User
	FullYears    int       `json:"full_years"`
	LastOnline   string    `json:"last_online"`
	Location     string    `json:"location"`
	IsBanned     bool      `json:"banned"`
	About        string    `json:"about"`
	AboutHTML    string    `json:"about_html"`
	CommonInfo   []string  `json:"common_info"`
	ShowComments bool      `json:"show_comments"`
	InFriends    bool      `json:"in_friends"`
	IsIgnored    bool      `json:"is_ignored"`
	Stats        UserStats `json:"stats"`
	StyleID      int       `json:"style_id"`
}

type UserStats struct {
	Statuses     UserStatuses        `json:"statuses"`
	FullStatuses UserStatuses        `json:"full_statuses"`
	Scores       UserScores          `json:"scores"`
	Types        UserTypes           `json:"types"`
	Ratings      UserRatings         `json:"ratings"`
	HasAnime     bool                `json:"has_anime?"`
	HasManga     bool                `json:"has_manga?"`
	Activity     []IntSliceNameValue `json:"activity"`
	Genres       []Genre             `json:"genres"`
	Studios      []Studio            `json:"studios"`
	Publishers   []Publisher         `json:"publishers"`
}

type UserStatuses struct {
	Anime []UserStatusesTitle `json:"anime"`
	Manga []UserStatusesTitle `json:"manga"`
}

type UserStatusesTitle struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	GroupedID string `json:"grouped_id"`
	Type      string `json:"type"`
}

type UserScores struct {
	Anime []StringNameValue `json:"anime"`
	Manga []StringNameValue `json:"manga"`
}

type UserTypes struct {
	Anime []StringNameValue `json:"anime"`
	Manga []StringNameValue `json:"manga"`
}

type UserRatings struct {
	Anime []StringNameValue `json:"anime"`
}

type UserRate struct {
	Rate
	Anime *Anime `json:"anime"`
	Manga *Manga `json:"manga"`
}

type UserFavourites struct {
	Animes     []UserFavourite `json:"animes"`
	Mangas     []UserFavourite `json:"mangas"`
	Characters []UserFavourite `json:"characters"`
	People     []UserFavourite `json:"people"`
	Mangakas   []UserFavourite `json:"mangakas"`
	Seyu       []UserFavourite `json:"seyu"`
	Producers  []UserFavourite `json:"producers"`
}

type UserFavourite struct {
	Base
	Image URL `json:"image"`
}

type UserUnread struct {
	Messages      int `json:"messages"`
	News          int `json:"news"`
	Notifications int `json:"notifications"`
}

type UserHistoryItem struct {
	ID          int         `json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	Description string      `json:"description"`
	Target      interface{} `json:"target"`
}

type UserBan struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Comment         Comment   `json:"comment"`
	ModeratorID     int       `json:"moderator_id"`
	Reason          string    `json:"reason"`
	CreatedAt       time.Time `json:"created_at"`
	DurationMinutes int       `json:"duration_minutes"`
	Moderator       User      `json:"moderator"`
}

// Users returns a list of users on specific page with specific count.
// You can use Search struct as request.
// Specify Query in request, if you want to filter results by nickname.
func (api *API) Users(req Request) ([]User, error) {
	var users []User
	if err := api.Raw(EndpointUsers, req.Values(), &users); err != nil {
		return nil, err
	}

	for i := range users {
		users[i].api = api
	}
	return users, nil
}

// UserByID returns an extended user info by its integer ID.
func (api *API) UserByID(id int) (UserExt, error) {
	var user UserExt
	user.api = api
	return user, api.Raw(EndpointUsers.With(id), nil, &user)
}

// UserByNickname returns an extended user by its nickname.
func (api *API) UserByNickname(name string) (UserExt, error) {
	var user UserExt
	user.api = api

	endpoint := EndpointUsers.With(name)
	return user, api.Raw(endpoint, Params{"is_nickname": 1}.Values(), &user)
}

// Whoami returns a brief information about current user. AccessToken required.
func (api *API) Whoami() (User, error) {
	if api.AccessToken == "" {
		return User{}, ErrAccessToken
	}

	var user User
	user.api = api
	return user, api.Raw(EndpointUsers.With("whoami"), nil, &user)
}

// Info returns a brief information about user.
func (user User) Info() (UserInfo, error) {
	var info UserInfo
	endpoint := EndpointUsers.With(user.ID, "info")
	return info, user.getAPI().Raw(endpoint, nil, &info)
}

// Friends returns user's friends.
func (user User) Friends() ([]User, error) {
	var users []User
	endpoint := EndpointUsers.With(user.ID, "friends")
	if err := user.getAPI().Raw(endpoint, nil, &users); err != nil {
		return nil, err
	}

	for i := range users {
		users[i].api = user.api
	}
	return users, nil
}

// Clubs returns user's clubs.
func (user User) Clubs() ([]Club, error) {
	var clubs []Club
	endpoint := EndpointUsers.With(user.ID, "clubs")
	return clubs, user.getAPI().Raw(endpoint, nil, &clubs)
}

// AnimeRates returns user's anime list.
func (user User) AnimeRates() ([]UserRate, error) {
	var rates []UserRate
	endpoint := EndpointUsers.With(user.ID, "anime_rates")
	return rates, user.getAPI().Raw(endpoint, nil, &rates)
}

// MangaRates returns user's manga list.
func (user User) MangaRates() ([]UserRate, error) {
	var rates []UserRate
	endpoint := EndpointUsers.With(user.ID, "manga_rates")
	return rates, user.getAPI().Raw(endpoint, nil, &rates)
}

// Favourites returns user's favourites.
func (user User) Favourites() (UserFavourites, error) {
	var favs UserFavourites
	endpoint := EndpointUsers.With(user.ID, "favourites")
	return favs, user.getAPI().Raw(endpoint, nil, &favs)
}

// Messages returns user's messages. AccessToken required.
// Kind must be one of: Inbox, Private, Sent, News, Notifications.
func (user User) Messages(kind string, req ...Request) ([]Message, error) {
	if len(req) == 0 {
		req = append(req, Params{})
	}

	values := req[0].Values()
	values.Set("type", kind)

	var msgs []Message
	endpoint := EndpointUsers.With(user.ID, "messages")
	if err := user.getAPI().Raw(endpoint, values, &msgs); err != nil {
		return nil, err
	}

	for i := range msgs {
		msgs[i].From.api, msgs[i].To.api = user.api, user.api
		msgs[i].Linked.setAPI(user.api)
	}
	return msgs, nil
}

// Unread retuns user's unread messages info. AccessToken required.
func (user User) Unread() (UserUnread, error) {
	var unread UserUnread
	endpoint := EndpointUsers.With(user.ID, "unread_messages")
	return unread, user.getAPI().Raw(endpoint, nil, &unread)
}

// History returns user history.
func (user User) History(search Search) ([]UserHistoryItem, error) {
	var history []UserHistoryItem
	endpoint := EndpointUsers.With(user.ID, "history")
	return history, user.getAPI().Raw(endpoint, nil, &history)
}

// Bans returns all user's bans.
func (user User) Bans() ([]UserBan, error) {
	var bans []UserBan
	endpoint := EndpointUsers.With(user.ID, "bans")
	if err := user.getAPI().Raw(endpoint, nil, &bans); err != nil {
		return nil, err
	}

	for i := range bans {
		bans[i].Moderator.api = user.api
	}
	return bans, nil
}
