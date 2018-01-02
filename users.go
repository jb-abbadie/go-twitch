package twitch

// UserList is the list of users outputted by GetStream
type UserList struct {
	Data []Stream `json:"data"`
}

// User is a twitch User
type User struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
}

// GetUserInput adds filters to the GetUser function
type GetUserInput struct {
	ID    string `url:"id,omitempty"`
	Login string `url:"login,omitempty"`
}

// GetUser returns a list of twitch user https://dev.twitch.tv/docs/api/reference#get-users
func (s Session) GetUser(input GetUserInput) (UserList, error) {
	var out UserList
	err := s.doRequest("/users", &input, &out)
	if err != nil {
		return UserList{}, err
	}
	return out, nil
}
