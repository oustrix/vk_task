package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type longPollServerResponse map[string]struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	TS     string `json:"ts"`
}

const (
	longPollServerRequest = "https://api.vk.com/method/groups.getLongPollServer"
	apiVersion            = "5.131"
)

func (b *Bot) Polling() error {
	s, err := b.getLongPollSession()
	if err != nil {
		return err
	}
	b.session = s

	return nil
}

func (b *Bot) getLongPollSession() (session, error) {
	params := url.Values{}
	params.Add("access_token", b.cfg.token)
	params.Add("v", apiVersion)
	params.Add("group_id", b.cfg.groupID)

	u, _ := url.ParseRequestURI(longPollServerRequest)
	u.RawQuery = params.Encode()
	url := fmt.Sprintf("%v", u)

	response, err := http.Get(url)
	if err != nil {
		return session{}, err
	}

	dec := json.NewDecoder(response.Body)

	var res longPollServerResponse
	err = dec.Decode(&res)
	if err != nil {
		return session{}, err
	}

	s := res["response"]
	ts, err := strconv.Atoi(s.TS)
	if err != nil {
		return session{}, errors.New("error while converting TS to integer")
	}

	return session{
		key:    s.Key,
		server: s.Server,
		ts:     ts,
	}, nil
}
