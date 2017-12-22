package twitch

import "testing"

func TestGetUserOptions(t *testing.T) {
	ts := emptyHTTPServer(t, "/users?login=test_user")
	defer ts.Close()

	testSession := Session{BaseURL: ts.URL}
	input := GetUserInput{Login: "test_user"}

	_, err := testSession.GetUser(input)
	if err != nil {
		t.Error(err)
	}
}
