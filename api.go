package goshiki

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

// New returns an API instance with default http client.
// AppName is your oauth2 application's name created on Shikimori site.
// Leave token as nil if you don't use functions that requires AccessToken.
// Visit https://shikimori.one/oauth for more information.
func New(appname string, token *oauth2.Token) *API {
	api := &API{Client: http.DefaultClient, AppName: appname}
	if token != nil {
		api.AccessToken = token.AccessToken
	}
	return api
}

// API allows you to interact with the Shikimori API.
type API struct {
	Client      *http.Client
	AppName     string
	AccessToken string
}

// Raw lets you make GET request manually. It also decodes JSON response,
// so you should pass pointer to save result.
func (api *API) Raw(endpoint Endpoint, values url.Values, v interface{}) error {
	url := BaseURL + string(endpoint)
	if values != nil {
		url += "?" + values.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	if api.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+api.AccessToken)
	}
	req.Header.Set("User-Agent", api.AppName)

	resp, err := api.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decodeJSON(resp.StatusCode, resp.Body, v)
}

func decodeJSON(code int, r io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if code == http.StatusNotFound {
		var e Error
		json.Unmarshal(data, &e)
		return e
	}

	if err = json.Unmarshal(data, v); err != nil {
		if _, ok := err.(*json.SyntaxError); ok {
			return errors.New(string(data))
		}
	}
	return err
}
