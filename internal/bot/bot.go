package bot

type Bot struct {
	token string
	cfg   *botConfig
}

func NewBot(token string) *Bot {
	botCfg := newConfig()
	botCfg.token = token

	return &Bot{
		cfg: botCfg,
	}
}
