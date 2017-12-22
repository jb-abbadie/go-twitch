package twitch

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/grafov/m3u8"
)

// StreamExtracter is a simple interface for ExtractStreamUrl
type StreamExtracter interface {
	ExtractStreamUrl(name string) ([]HLSStream, error)
}

// ExtractStreamUrl returns the lists of urls for each stream quality
func (s Session) ExtractStreamUrl(name string) ([]HLSStream, error) {

	data, _ := s.getHLSAccessToken(name, "https://api.twitch.tv/api/channels/")
	playlist := getChannelM3U8Playlist(name, data, "https://usher.ttvnw.net/api/channel/hls/")
	pl := parsePlaylist(playlist)

	return pl, nil
}

// HLSStream shows information about a stream, URL is curl to an M3U6 playlist
type HLSStream struct {
	URL        string
	Quality    string
	Resolution string
	Bitrate    uint32
}

type accessToken struct {
	MobileRestricted bool   `json:"mobile_restricted"`
	Sig              string `json:"sig"`
	Token            string `json:"token"`
}

func (s Session) getHLSAccessToken(channel string, url string) (accessToken, error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", url+channel+"/access_token", nil)
	if err != nil {
		return accessToken{}, err
	}
	req.Header.Add("Client-ID", s.ClientID)

	resp, err := client.Do(req)
	if err != nil {
		return accessToken{}, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return accessToken{}, err
	}
	var data accessToken
	err = json.Unmarshal(out, &data)
	return data, err
}

type getChannelOptions struct {
	P           int    `url:"p"`
	AllowSource bool   `url:"allow_source"`
	Token       string `url:"token"`
	Sig         string `url:"sig"`
}

func getChannelM3U8Playlist(channel string, at accessToken, url string) io.Reader {
	client := new(http.Client)

	options := getChannelOptions{
		P:           rand.Intn(999999),
		AllowSource: true,
		Token:       at.Token,
		Sig:         at.Sig,
	}
	v, _ := query.Values(options)

	reqUrl := url + channel + ".m3u8?" + v.Encode()
	req, _ := http.NewRequest("GET", reqUrl, nil)
	ret, _ := client.Do(req)
	return ret.Body
}

func parsePlaylist(pl io.Reader) []HLSStream {
	p, _, err := m3u8.DecodeFrom(pl, true)
	if err != nil {
		panic(err)
	}

	variants := p.(*m3u8.MasterPlaylist).Variants
	out := make([]HLSStream, len(variants))
	for i := 0; i < len(variants); i++ {
		out[i] = HLSStream{
			variants[i].URI,
			variants[i].Video,
			variants[i].Resolution,
			variants[i].Bandwidth,
		}
	}
	return out
}
