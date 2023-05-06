package botapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type longPollDetails struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	TS     string `json:"ts"`
}

const (
	longPollServerRequest = "https://api.vk.com/method/groups.getLongPollServer"
	apiVersion            = "5.131"
)

// Polling is main polling function
func (b *Bot) Polling() error {
	details, err := b.getLongPollSession()
	if err != nil {
		return err
	}

	b.cfg.pollConfig = newLongPollConfig(
		details.Server, details.Key, details.TS)

	return nil
}

// getLongPollSession is for get session data
func (b *Bot) getLongPollSession() (*longPollDetails, error) {
	params := url.Values{}
	params.Add("access_token", b.cfg.token)
	params.Add("v", apiVersion)
	params.Add("group_id", b.cfg.groupID)

	u, _ := url.ParseRequestURI(longPollServerRequest)
	u.RawQuery = params.Encode()
	url := fmt.Sprintf("%v", u)

	response, err := http.Get(url)
	if err != nil {
		return &longPollDetails{}, err
	}

	dec := json.NewDecoder(response.Body)

	var res map[string]longPollDetails
	err = dec.Decode(&res)
	if err != nil {
		return &longPollDetails{}, err
	}

	details, ok := res["response"]
	if !ok {
		return &longPollDetails{}, errors.New("there is no \"response\" field in server response")
	}

	return &details, nil
}
