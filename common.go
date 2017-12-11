package twitch

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/go-querystring/query"
)

type TwitchInterface interface {
	GetStream(input GetStreamInput) (StreamList, error)
}

// Session is a base struct for calling the other functions
type Session struct {
	BaseURL  string
	ClientID string
}

// NewSession creates a new Twitch Session
func NewSession() Session {
	client_id := os.Getenv("CLIENT_ID")
	return Session{"https://api.twitch.tv/helix", client_id}
}

func (s *Session) doRequest(path string, q interface{}, r interface{}) error {

	client := new(http.Client)
	reqURL, err := buildURL(s.BaseURL, path, q)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Client-ID", s.ClientID)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Failed request " + resp.Status)
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(out, &r)
	return err
}

func buildURL(baseURL string, path string, r interface{}) (string, error) {
	outURL := baseURL + path

	v, err := query.Values(r)
	if err != nil {
		return "", err
	}
	qString := v.Encode()
	if qString == "" {
		return outURL, nil
	}
	return outURL + "?" + qString, nil
}
