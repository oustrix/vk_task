package botapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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

// InitSession is using for init session.
func (b *Bot) InitSession() error {
	details, err := b.getLongPollSession()
	if err != nil {
		return err
	}

	ts, err := strconv.Atoi(details.TS)
	if err != nil {
		return errors.New("error while converting TS to integer")
	}

	b.PollConfig = newLongPollConfig(
		details.Server, details.Key, ts)

	return nil
}

// getLongPollSession is for get session data.
func (b *Bot) getLongPollSession() (*longPollDetails, error) {
	params := url.Values{}
	params.Add("access_token", b.token)
	params.Add("v", apiVersion)
	params.Add("group_id", b.GroupID)

	u, _ := url.ParseRequestURI(longPollServerRequest)
	u.RawQuery = params.Encode()
	url := fmt.Sprintf("%v", u)

	response, err := http.Get(url) //nolint:gosec
	if err != nil {
		return &longPollDetails{}, err
	}
	defer response.Body.Close()

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
