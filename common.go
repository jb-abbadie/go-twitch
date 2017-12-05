package twitch

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"os"
)

type Session struct {
	BaseURL  string
	ClientID string
}

func NewSession() Session {
	client_id := os.Getenv("CLIENT_ID")
	return Session{"https://api.twitch.tv/helix", client_id}
}

func (s *Session) doRequest(path string, q interface{}, r interface{}) error {

	client := new(http.Client)
	reqURL, _ := buildURL(s.BaseURL, path, q)
	req, _ := http.NewRequest("GET", reqURL, nil)

	req.Header.Add("Client-ID", s.ClientID)

	resp, _ := client.Do(req)
	out, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(out, &r)
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
