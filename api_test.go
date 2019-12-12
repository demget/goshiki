package goshiki_test

import (
	"os"
	"testing"

	"github.com/demget/goshiki"
)

var (
	api        = goshiki.New("goshiki", nil)
	apiNoToken = goshiki.New("goshiki", nil)
	nickname   string
)

func TestMain(m *testing.M) {
	api.AccessToken, nickname =
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("NICKNAME")

	os.Exit(m.Run())
}
