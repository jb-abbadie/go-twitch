package twitch

import (
	"testing"
)

func TestImplementInterface(t *testing.T) {
	var twitch TwitchInterface = (*Session)(nil)
	t.Log(twitch)
}
