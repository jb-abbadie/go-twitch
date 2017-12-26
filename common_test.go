package twitch

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	s := NewSession("foobar")

	if s.ClientID != "foobar" {
		t.Fatal("Failed creating session")
	}

}

func TestImplementInterface(t *testing.T) {
	assert.Implements(t, (*Interface)(nil), new(Session))
}

type readf func(p []byte) (n int, err error)

type FakeReader struct {
	readFunc readf
}

func (f FakeReader) Read(p []byte) (n int, err error) {
	return f.readFunc(p)
}

func TestParseJSONFails(t *testing.T) {
	fr := FakeReader{func(p []byte) (n int, err error) {
		return 0, errors.New("Test")
	},
	}

	err := parseJSON(fr, nil)
	if err == nil {
		t.Fail()
	}
}

func TestBuildURLFail(t *testing.T) {
	_, err := buildURL("", "", 8)
	if err == nil {
		t.Fail()
	}
}

func TestBuildTwitchReqFailInvalidQ(t *testing.T) {
	s := Session{}
	_, err := s.buildTwitchReq("", "", 8)
	assert.Contains(t, err.Error(), "expects struct input. Got int")
}

func TestBuildTwitchReqFailDuringHTTPreq(t *testing.T) {
	s := Session{}
	_, err := s.buildTwitchReq("", "%zzzz", GetUserInput{})
	assert.Contains(t, err.Error(), "parse %zzzz: invalid URL escape ")
}

func TestBuildDoRequestFailBuild(t *testing.T) {
	s := Session{}
	err := s.doRequest("%zzzz", nil, nil)
	assert.Contains(t, err.Error(), "parse")
}

func TestDoRequestInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{")
	}))
	defer ts.Close()
	s := Session{}
	err := s.doRequest(ts.URL, nil, nil)
	assert.Contains(t, err.Error(), "JSON input")
}

func TestDoRequestInvalidResponseCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(120)
	}))
	defer ts.Close()
	s := Session{}
	err := s.doRequest(ts.URL, nil, nil)
	assert.Contains(t, err.Error(), "status code 120")
}

func TestDoRequestInvalidHTTPResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(301)
	}))
	defer ts.Close()
	s := Session{}
	err := s.doRequest(ts.URL, nil, nil)
	assert.Contains(t, err.Error(), "301 response missing Location header")
}
