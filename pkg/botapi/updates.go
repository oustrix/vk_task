package botapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Update struct {
	Type   string                 `json:"type"`
	Object map[string]interface{} `json:"object"`
}

type Message struct {
	ID     int    `json:"id"`
	Date   int    `json:"date"`
	PeerID int    `json:"peer_id"`
	FromID int    `json:"from_id"`
	Text   string `json:"text"`
}

type UpdatesChannel <-chan Update

func (b *Bot) GetUpdates() ([]Update, error) {
	params := url.Values{}
	params.Add("act", "a_check")
	params.Add("key", b.PollConfig.key)
	params.Add("ts", fmt.Sprintf("%d", b.PollConfig.ts))
	params.Add("wait", fmt.Sprintf("%d", b.PollConfig.Wait))

	u, _ := url.ParseRequestURI(b.PollConfig.server)
	u.RawQuery = params.Encode()
	url := fmt.Sprintf("%v", u)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(res.Body)

	var resBody map[string]interface{}
	err = dec.Decode(&resBody)
	if err != nil {
		return nil, err
	}

	if resBody["failed"] != nil {
		if resBody["failed"].(float64) == 1 {
			b.PollConfig.ts = int(resBody["ts"].(float64))
		} else {
			err := b.InitSession()
			if err != nil {
				return nil, err
			}
		}

		return nil, errors.New("failed to get updates. code: " + fmt.Sprintf("%v", resBody["failed"].(float64)))
	}

	b.PollConfig.ts, err = strconv.Atoi(resBody["ts"].(string))
	if err != nil {
		return nil, errors.New("error while converting TS to integer")
	}

	var updates []Update
	for _, update := range resBody["updates"].([]interface{}) {
		updates = append(updates, Update{
			Type:   update.(map[string]interface{})["type"].(string),
			Object: update.(map[string]interface{})["object"].(map[string]interface{}),
		})
	}

	return updates, nil
}

func (b *Bot) GetUpdatesChan() UpdatesChannel {
	ch := make(chan Update, b.Buffer)

	go func() {
		for {
			select {
			case <-b.shutdownChannel:
				close(ch)
				return
			default:
			}

			updates, err := b.GetUpdates()
			if err != nil {
				log.Println(err)
				log.Println("Failed to get updates, retrying in 5 seconds...")
				time.Sleep(5 * time.Second)

				continue
			}

			for _, update := range updates {
				ch <- update
			}
		}
	}()

	return ch
}

func (b *Bot) StopReceivingUpdates() {
	close(b.shutdownChannel)
}
