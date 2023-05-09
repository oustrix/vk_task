package botapi

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

var SendMessageErrorCodes = map[int]string{
	104: "Not found",
	900: "Can't send messages for users from blacklist",
	901: "Can't send messages for users without permission",
	902: "Can't send messages to this user due to their privacy settings",
	911: "Keyboard format is invalid",
	912: "This is a chat bot feature, change this status in settings",
	913: "Too many forwarded messages",
	914: "Message is too long",
	917: "You don't have access to this chat",
	921: "Can't forward these messages",
	922: "You left this chat",
	925: "You are not admin of this chat",
	936: "Contact not found940Too many posts in messages",
	943: "Cannot use this intent",
	944: "Limits overflow for this intent",
	945: "Chat was disabled",
	946: "Chat not supported",
	950: "Can't send message, reply timed out",
	962: "You can't access donut chat without subscription",
	969: "Message cannot be forwarded",
	979: "App action is restricted for conversations with communities",
}

type MessageConfig struct {
	UserID          int             `url:"user_id"`
	RandomID        int             `url:"random_id"`
	PeerID          int             `url:"peer_id,omitempty"`
	PeerIDs         int             `url:"peer_ids,omitempty"`
	Domain          int             `url:"domain,omitempty"`
	ChatID          int             `url:"chat_id,omitempty"`
	UserIDs         int             `url:"user_ids,omitempty"`
	Message         string          `url:"message,omitempty"` // required if no attachment, max length - 4096
	Guid            string          `url:"guid,omitempty"`
	Lat             int             `url:"lat,omitempty"`
	Long            int             `url:"long,omitempty"`
	Attachment      string          `url:"attachment,omitempty"` // required if no message
	ReplyTo         int             `url:"reply_to,omitempty"`
	ForwardMessages string          `url:"forward_messages,omitempty"`
	Forward         ForwardMessages `url:"forward,omitempty"`
	StickerID       int             `url:"sticker_id,omitempty"`
	GroupID         int             `url:"group_id,omitempty"`
	Keyboard        Keyboard        `url:"keyboard,omitempty"`
	Template        MessageTemplate `url:"template,omitempty"`
	Payload         string          `url:"payload,omitempty"`
	ContentSource   string          `url:"content_source,omitempty"`
	DontParseLinks  int             `url:"dont_parse_links,omitempty"`
	DisableMentions int             `url:"disable_mentions,omitempty"`
	Intent          string          `url:"intent,omitempty"`
	SubscribeID     int             `url:"subscribe_id,omitempty"`
}

type ForwardMessages struct {
	OwnerID                int   `url:"owner_id,omitempty"`
	PeerID                 int   `url:"peer_id,omitempty"`
	ConversationMessageIDs []int `url:"conversation_message_ids,omitempty"`
	MessageIDs             []int `url:"message_ids,omitempty"`
	IsReply                bool  `url:"is_reply,omitempty"`
}

type MessageTemplate struct {
	Title       string                 `url:"title,omitempty"`
	Description string                 `url:"description,omitempty"`
	PhotoID     string                 `url:"photo_id,omitempty"`
	Buttons     []Button               `url:"buttons,omitempty"`
	Action      map[string]interface{} `url:"action,omitempty"`
}

func (b *Bot) SendMessage(config *MessageConfig) error {
	if config.Message == "" && config.Attachment == "" {
		return fmt.Errorf("message or attachment required")
	}

	params, err := query.Values(*config)
	if err != nil {
		return err
	}

	res, err := b.CallMethod("messages.send", params.Encode())
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("status code %d", res.StatusCode)
	}

	dec := json.NewDecoder(res.Body)

	var resBody map[string]interface{}
	err = dec.Decode(&resBody)
	if err != nil {
		return err
	}

	if resBody["error"] != nil {
		errCode := int(resBody["error"].(map[string]interface{})["error_code"].(float64))
		return fmt.Errorf("error code %d: %s", errCode, SendMessageErrorCodes[errCode])
	}

	return nil
}

func (b *Bot) CallMethod(method string, params string) (*http.Response, error) {
	apiURL := "https://api.vk.com/method/"

	endpoint := apiURL + method + "?" + params + "&access_token=" + b.token + "&v" + apiVersion

	return http.Get(endpoint)
}
