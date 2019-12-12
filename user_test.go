package goshiki_test

import (
	"testing"
	"time"

	"github.com/demget/goshiki"
	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	defer time.Sleep(time.Second)

	users, err := api.Users(goshiki.Search{Page: 1, Limit: 3})
	assert.NoError(t, err)
	assert.Len(t, users, 4)

	for _, user := range users {
		info, err := user.Info()
		assert.NoError(t, err)
		assert.NotEmpty(t, info.Locale)
	}
}

func TestWhoami(t *testing.T) {
	defer time.Sleep(time.Second)

	_, err := apiNoToken.Whoami()
	assert.Error(t, err)

	user, err := api.Whoami()
	assert.NoError(t, err)
	assert.Equal(t, nickname, user.Nickname)

	// Now you need messages scopes for this:
	//
	// msgs, err := user.Messages(goshiki.News)
	// assert.NoError(t, err)
	// assert.NotEmpty(t, msgs) // You must have unchecked news to complete test

	// unread, err := user.Unread()
	// assert.NoError(t, err)
	// assert.Zero(t, unread.Messages) // You must read all your inbox messages
}

func TestUserBy(t *testing.T) {
	defer time.Sleep(time.Second)

	user, err := api.UserByID(999999)
	assert.Error(t, err)
	assert.Equal(t, user.ID, 0)

	user1, err := api.UserByID(1)
	assert.NoError(t, err)

	user2, err := api.UserByNickname("morr")
	assert.NoError(t, err)

	assert.Equal(t, user1, user2)
}

func _TestUserInfo(t *testing.T) {
	defer time.Sleep(time.Second)

	user, err := api.UserByNickname("dox5")
	assert.NoError(t, err)

	info, err := user.Info()
	assert.NoError(t, err)
	assert.NotNil(t, info.BirthOn)

	time := time.Time(*info.BirthOn)
	assert.Equal(t, "1930-01-01", time.Format("2006-01-02"))
}

func TestUser(t *testing.T) {
	defer time.Sleep(time.Second)

	assert.Panics(t, func() {
		goshiki.User{ID: 1}.Info()
	})

	user, err := api.UserByNickname("morr")
	assert.NoError(t, err)

	friends, err := user.Friends()
	assert.NoError(t, err)
	assert.NotEmpty(t, friends)

	clubs, err := user.Clubs()
	assert.NoError(t, err)
	assert.NotEmpty(t, clubs)

	ars, err := user.AnimeRates()
	assert.NoError(t, err)
	assert.NotEmpty(t, ars)

	mrs, err := user.MangaRates()
	assert.NoError(t, err)
	assert.NotEmpty(t, mrs)

	time.Sleep(time.Second)

	favs, err := user.Favourites()
	assert.NoError(t, err)
	assert.NotEmpty(t, favs.Animes)

	history, err := user.History(goshiki.Search{})
	assert.NoError(t, err)
	assert.NotEmpty(t, history)

	bans, err := user.Bans()
	assert.NoError(t, err)
	assert.NotEmpty(t, bans)

	info, err := bans[0].Moderator.Info()
	assert.NoError(t, err)
	assert.NotEmpty(t, info.Locale)
}
