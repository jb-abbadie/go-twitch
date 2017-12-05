package twitch

import (
	"time"
)

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

func (s *Session) GetStream(input *GetStreamInput) (*StreamList, error) {
	var out StreamList
	err := s.doRequest("/streams", input, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
