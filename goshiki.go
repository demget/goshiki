package goshiki

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	BaseURL = "https://shikimori.one/api"

	EndpointAchievements Endpoint = "/achievements"
	EndpointAnimeVideos  Endpoint = "/anime_videos"
	EndpointAnimes       Endpoint = "/animes"
	EndpointAppear       Endpoint = "/appear"
	EndpointBans         Endpoint = "/bans"
	EndpointCalendars    Endpoint = "/calendars"
	EndpointCharacters   Endpoint = "/characters"
	EndpointClubs        Endpoint = "/clubs"
	EndpointComments     Endpoint = "/comments"
	EndpointDevices      Endpoint = "/devices"
	EndpointDialogs      Endpoint = "/dialogs"
	EndpointForums       Endpoint = "/forums"
	EndpointFriends      Endpoint = "/friends"
	EndpointGenres       Endpoint = "/genres"
	EndpointMangas       Endpoint = "/mangas"
	EndpointMessages     Endpoint = "/messages"
	EndpointPeople       Endpoint = "/people"
	EndpointPublishers   Endpoint = "/publishers"
	EndpointRanobe       Endpoint = "/ranobe"
	EndpointSessions     Endpoint = "/sessions"
	EndpointStats        Endpoint = "/stats"
	EndpointStudios      Endpoint = "/studios"
	EndpointStyles       Endpoint = "/styles"
	EndpointTopics       Endpoint = "/topics"
	EndpointUserImages   Endpoint = "/user_images"
	EndpointUserRates    Endpoint = "/user_rates"
	EndpointUsers        Endpoint = "/users"
	EndpointVideos       Endpoint = "/videos"
)

// ErrAccessToken is returned when AccessToken is empty, but function
// requires it to do request.
var ErrAccessToken = errors.New("goshiki: AccessToken is empty")

// Endpoint represents Shikimori API endpoint.
type Endpoint string

// With returns endpoint with passed values separated by '/'.
func (e Endpoint) With(v ...interface{}) Endpoint {
	for i := range v {
		e += Endpoint("/" + fmt.Sprint(v[i]))
	}
	return e
}

// Base represents basic fields for most of the Shikimori objects.
type Base struct {
	ID     int    `json:"id"`
	NameEn string `json:"name"`
	NameRu string `json:"russian"`
	Image  Image  `json:"image"`
	URL    URL    `json:"url"`
}

// StringNameValue represents a common key-value pair in Shikimori API
// where Name is string.
type StringNameValue struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// IntNameValue represents a common key-value pair in Shikimori API
// where Name is integer.
type IntNameValue struct {
	Name  int `json:"name"`
	Value int `json:"value"`
}

// IntSliceNameValue represents a common key-value pair in Shikimori API
// where Name is slice of integers.
type IntSliceNameValue struct {
	Name  []int `json:"name"`
	Value int   `json:"value"`
}

// Error represents Shikimori API error with message and code.
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Error implements error interface.
func (e Error) Error() string {
	return fmt.Sprintf("goshiki: %s (%d)", e.Message, e.Code)
}

// URL represents Shikimori path without scheme and host.
type URL string

// Full returns full url of Shikimori path.
func (url URL) Full() string {
	return "https://shikimori.org" + string(url)
}

// Image represents Shikimori image with possible sizes.
type Image struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
	X96      string `json:"x96"`
	X48      string `json:"x48"`

	X160 string `json:"x160"` // User only
	X148 string `json:"x148"` // User only
	X80  string `json:"x80"`  // User only
	X64  string `json:"x64"`  // User only
	X32  string `json:"x32"`  // User only
	X16  string `json:"x16"`  // User only
}

// Time represents a time for custom JSON unmarshaling.
type Time time.Time

// UnmarshalJSON converts "yyyy-MM-dd" pattern to valid Time instance.
func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}

	parsed, err := time.Parse("2006-01-02", s)
	*t = Time(parsed)
	return err
}
