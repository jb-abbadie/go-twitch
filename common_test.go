package twitch

import (
	"os"
	"testing"
)

func TestNewSession(t *testing.T) {
	old_env := os.Getenv("CLIENT_ID")
	os.Setenv("CLIENT_ID", "foobar")
	defer os.Setenv("CLIENT_ID", old_env)

	s := NewSession()
	if s.ClientID != "foobar" {
		t.Fatal("Failed creating session")
	}

}

func TestImplementInterface(t *testing.T) {
	var twitch TwitchInterface = (*Session)(nil)
	t.Log(twitch)
}
