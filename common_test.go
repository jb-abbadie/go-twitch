package twitch

import (
	"errors"
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
	var twitch Interface = (*Session)(nil)
	t.Log(twitch)
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
