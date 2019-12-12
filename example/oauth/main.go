package main

import (
	"fmt"

	"github.com/demget/goshiki"
	"github.com/pkg/browser"
)

const (
	id     = "8f6f3e9340d313175a9adda0f8c119c30267272be615ddda22e62c69b7c772c5"
	secret = "6fb865b4ad829b7faa896cc6850d5f026b270d8e79f10945aa28719f994d262a"
	uri    = "urn:ietf:wg:oauth:2.0:oob"
)

func main() {
	// Creating a new Auth instance
	auth := goshiki.Auth(id, secret, uri)

	if err := browser.OpenURL(auth.URL()); err != nil {
		panic(err)
	}

	fmt.Print(">> ")
	var code string
	fmt.Scanln(&code)

	token, err := auth.Token(code)
	if err != nil {
		panic(err)
	}

	// Creating a new API intsance
	api := goshiki.New("goshiki", token)

	user, err := api.Whoami()
	if err != nil {
		panic(err)
	}

	fmt.Println("AccessToken:", token.AccessToken)
	fmt.Println("Nickname:", user.Nickname)
}
