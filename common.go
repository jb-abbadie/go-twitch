package twitch

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Interface for the whole lib
type Interface interface {
	GetStreamer
	StreamExtracter
}

// Session is a base struct for calling the other functions
type Session struct {
	BaseURL  string
	ClientID string
}

// NewSession creates a new Twitch Session
func NewSession(clientID string) Session {
	return Session{"https://api.twitch.tv/helix", clientID}
}

func (s *Session) doRequest(path string, q interface{}, r interface{}) error {

	req, err := s.buildTwitchReq("GET", path, q)
	if err != nil {
		return err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Failed request " + resp.Status)
	}
	return parseJSON(resp.Body, r)
}

func (s *Session) buildTwitchReq(method string, path string, q interface{}) (*http.Request, error) {
	reqURL, err := buildURL(s.BaseURL, path, q)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Client-ID", s.ClientID)
	return req, nil
}

func parseJSON(resp io.Reader, r interface{}) error {
	out, err := ioutil.ReadAll(resp)
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
