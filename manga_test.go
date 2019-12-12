package goshiki_test

import (
	"testing"
	"time"

	"github.com/demget/goshiki"
	"github.com/stretchr/testify/assert"
)

func TestMangas(t *testing.T) {
	defer time.Sleep(time.Second)

	mangas, err := api.Mangas(goshiki.Search{Query: "Bleach"},
		goshiki.ByChapters,     // order=chapters
		goshiki.WithManga,      // ...
		goshiki.NotOneshot,     // kind=manga,!one_shot
		goshiki.Publishers{83}) // publisher=83
	assert.NoError(t, err)
	assert.Len(t, mangas, 1)
}

func TestMangaBy(t *testing.T) {
	defer time.Sleep(time.Second)

	manga, err := api.MangaByID(12)
	assert.NoError(t, err)
	assert.Len(t, manga.RatesScoresStats, 10)

	time := time.Time(manga.AiredOn)
	assert.Equal(t, "2001-08-07", time.Format("2006-01-02"))
}
