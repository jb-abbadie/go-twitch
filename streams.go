package twitch

import (
	"time"
)

// StreamList is the list of stream outputted by GetStream
type StreamList struct {
	Data []Stream `json:"data"`
}

type Stream struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	GameID       string    `json:"game_id"`
	CommunityIds []string  `json:"community_ids"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
	Language     string    `json:"language"`
	ThumbnailURL string    `json:"thumbnail_url"`
}

// GetStreamInput adds filters to the GetStream function see GetStream for more information
type GetStreamInput struct {
	After       string `url:"after,omitempty"`
	Before      string `url:"before,omitempty"`
	CommunityId string `url:"community_id,omitempty"`
	First       int    `url:"first,omitempty"`
	GameID      string `url:"game_id,omitempty"`
	Language    string `url:"language,omitempty"`
	Type        string `url:"type,omitempty"`
	UserID      string `url:"user_id,omitempty"`
	UserLogin   string `url:"user_login,omitempty"`
}

// GetStreamer is a simple interface for the GetStream function
type GetStreamer interface {
	GetStream(input GetStreamInput) (StreamList, error)
}

/* GetStream returns a list of Streams
reference : https://dev.twitch.tv/docs/api/reference#get-streams */
func (s Session) GetStream(input GetStreamInput) (StreamList, error) {
	var out StreamList
	err := s.doRequest("/streams", &input, &out)
	if err != nil {
		return StreamList{}, err
	}
	return out, nil
}
