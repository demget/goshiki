package goshiki_test

import (
	"testing"
	"time"

	"github.com/demget/goshiki"
	"github.com/stretchr/testify/assert"
)

func TestRanobes(t *testing.T) {
	defer time.Sleep(time.Second)

	ranobes, err := api.Ranobes(goshiki.Search{Limit: 15},
		goshiki.Released,        // status=released
		goshiki.Season("!201x")) // season=!201x
	assert.NoError(t, err)
	assert.Len(t, ranobes, 15)
}

func TestRanobeBy(t *testing.T) {
	defer time.Sleep(time.Second)

	ranobe, err := api.RanobeByID(40171)
	assert.NoError(t, err)
	assert.Len(t, ranobe.RatesScoresStats, 10)

	time := time.Time(ranobe.AiredOn)
	assert.Equal(t, "2011-03-23", time.Format("2006-01-02"))
}
