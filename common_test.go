package twitch

import (
	"testing"
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
