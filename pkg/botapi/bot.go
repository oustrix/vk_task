package botapi

const (
	_BotBufferDefault = 100
)

type Bot struct {
	token      string
	GroupID    string
	Buffer     int
	PollConfig *longPollConfig
}

func NewBot(token string, groupID string) *Bot {
	return &Bot{
		token:   token,
		GroupID: groupID,
		Buffer:  _BotBufferDefault,
	}
}

func (b *Bot) SetTimeout(timeout int) {
	b.PollConfig.wait = timeout
}
