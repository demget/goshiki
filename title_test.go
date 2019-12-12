package goshiki_test

import (
	"testing"
	"time"

	"github.com/demget/goshiki"
	"github.com/stretchr/testify/assert"
)

func TestTitle(t *testing.T) {
	defer time.Sleep(time.Second)

	assert.Panics(t, func() {
		goshiki.Title{ID: 1}.Roles()
	})

	time.Sleep(time.Second)

	title, err := api.AnimeByID(21)
	assert.NoError(t, err)

	roles, err := title.Roles()
	assert.NoError(t, err)
	assert.NotEmpty(t, roles)

	similar, err := title.Similar()
	assert.NoError(t, err)
	assert.NotEmpty(t, similar)

	related, err := title.Related()
	assert.NoError(t, err)
	assert.NotEmpty(t, related)

	screens, err := title.AllScreenshots()
	assert.NoError(t, err)
	assert.NotEmpty(t, screens)

	time.Sleep(time.Second)

	franchise, err := title.Franchise()
	assert.NoError(t, err)
	assert.NotEmpty(t, franchise.Links)
	assert.NotEmpty(t, franchise.Nodes)

	links, err := title.ExternalLinks()
	assert.NoError(t, err)
	assert.NotEmpty(t, links)

	topics, err := title.Topics(goshiki.Topics{Kind: "episode"})
	assert.NoError(t, err)
	assert.NotEmpty(t, topics)
	assert.Equal(t, topics[0].Event, "episode")
}
