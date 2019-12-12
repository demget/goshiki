# goshiki

Goshiki is a Go package that provides easy interaction with [Shikimori API](https://shikimori.org/api/doc).

## Installation
```bash
$ go get -u github.com/demget/goshiki
```

## Usage
```go
// Creating a new API instance without token
api := goshiki.New("appname", nil)

user, err := api.UserByNickname("morr")
if err != nil {
    log.Fatal(err)
}

// User fields
user.ID
user.LastOnlineAt
user.Image.X160

// User functions
user.Friends() // api/users/morr/friends
user.Clubs()   // api/users/morr/clubs
user.Bans()    // api/users/morr/bans
```

#### Searching
```go
api.Animes(goshiki.Search{Query: "SAO", Limit: 7},
    goshiki.ByAiredOn,      // order=aired_on
    goshiki.NotOVA,         // ...
    goshiki.NotMovie,       // kind=!ova,!movie
    goshiki.Released,       // status=released
    goshiki.Season("201x"), // season=201x
    goshiki.DurationD)      // duration=D
```

#### Authorization
```go
// Creating a new Auth instance
auth := goshiki.Auth(id, secret, uri)

url := auth.URL() // authorize URL to get the code
code := "..."     // pushed code to RedirectURI

token, err := auth.Token(code)
if err != nil {
    log.Fatal(err)
}

// API with valid AccessToken
api := goshiki.New("goshiki", token) 

user, err := api.Whoami()
if err != nil {
    log.Fatal(err)
}

unread, err := user.Unread() // requires AccessToken
if err != nil {
    log.Fatal(err)
}

// Unread messages
unread.Messages
unread.News
unread.Notifications
```

You can also use `Raw` function for methods that not implemented yet.
```go
var msg goshiki.Message
api.Raw(goshiki.EndpointMessages.With(8), nil, &msg) // api/messages/8
```

## Implemented methods
| Endpoint       | Status | Endpoint | Status |
| :------------: | :----: | :------: | :----: |
| `achievements` | ❌ | `publishers`     | ❌ |
| `ranobe`       | ✅ | `signup`         | ❌ |
| `animes`       | ✅ | `sessions`       | ❌ |
| `appear`       | ❌ | `stats`          | ❌ |
| `bans`         | ❌ | `studios`        | ❌ |
| `calendars`    | ❌ | `styles`         | ❌ |
| `characters`   | ❌ | `topics`         | ❌ |
| `clubs`        | ❌ | `user_images`    | ❌ |
| `comments`     | ❌ | `user_rates`     | ❌ |
| `devices`      | ❌ | `users`          | ✅ |
| `dialogs`      | ❌ | `mangas`         | ✅ |
| `forums`       | ❌ | `people`         | ❌ |
| `friends`      | ❌ | `messages`       | ❌ |
| `genres`       | ❌ | `abuse_requests` | ❌ |