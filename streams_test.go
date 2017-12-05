package twitch

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetStreamsOptions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.RequestURI, "/streams?after=125")
		fmt.Fprintln(w, "{}")
	}))

	defer ts.Close()

	testSession := Session{ts.URL, ""}
	input := GetStreamInput{}
	input.After = "125"

	_, err := testSession.GetStream(&input)
	if err != nil {
		t.Error("error parsing stream", err)
	}
}

func TestGetStreamsHTTPrequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.RequestURI, "/streams")
		fmt.Fprintln(w, "{}")
	}))

	defer ts.Close()
	testSession := Session{ts.URL, ""}
	_, err := testSession.GetStream(nil)
	if err != nil {
		t.Error("error parsing stream", err)
	}
}

func TestGetStreamsParseJSON(t *testing.T) {
	fakeStream := StreamList{
		[]Stream{{
			"26901632320",
			"20786541",
			"488191",
			[]string{},
			"live",
			"HatFilms make an entire album!",
			41971,
			time.Unix(0, 0),
			"en",
			"https://static-cdn.jtvnw.net/previews-ttv/live_user_yogscast-{width}x{height}.jpg"}}}
	fakeJson, err := json.Marshal(fakeStream)
	if err != nil {
		t.Fatal("error creating json")
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(fakeJson))
	}))
	defer ts.Close()

	testSession := Session{ts.URL, ""}
	data, err := testSession.GetStream(nil)
	if err != nil {
		t.Error("error parsing stream")
	}
	assert.Equal(t, *data, fakeStream)
}
