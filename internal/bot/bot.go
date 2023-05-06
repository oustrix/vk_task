package bot

type Bot struct {
	token   string
	cfg     *botConfig
	session session
}

type session struct {
	key    string
	server string
	ts     int
}

func NewBot(token string, groupID string) *Bot {
	botCfg := newConfig()
	botCfg.token = token
	botCfg.groupID = groupID

	return &Bot{
		cfg: botCfg,
	}
}
