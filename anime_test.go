package goshiki_test

import (
	"testing"
	"time"

	"github.com/demget/goshiki"
	"github.com/stretchr/testify/assert"
)

func TestAnimes(t *testing.T) {
	defer time.Sleep(time.Second)

	animes, err := api.Animes(goshiki.Search{Query: "SAO", Limit: 7},
		goshiki.ByAiredOn,      // order=aired_on
		goshiki.NotOVA,         // ...
		goshiki.NotMovie,       // kind=!ova,!movie
		goshiki.Released,       // status=released
		goshiki.Season("201x"), // season=201x
		goshiki.Score(5),       // score=5
		goshiki.DurationD,      // duration=D
		goshiki.Censored)       // censored=true
	assert.NoError(t, err)
	assert.Len(t, animes, 7)
}

func TestAnimeBy(t *testing.T) {
	defer time.Sleep(time.Second)

	anime, err := api.AnimeByID(37999)
	assert.NoError(t, err)
	assert.Len(t, anime.RatesScoresStats, 10)

	time := time.Time(anime.AiredOn)
	assert.Equal(t, "2019-01-12", time.Format("2006-01-02"))
}
