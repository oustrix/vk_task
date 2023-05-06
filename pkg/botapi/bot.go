package botapi

type Bot struct {
	cfg *botConfig
}

func NewBot(token string, groupID string) *Bot {
	botCfg := newBotConfig()
	botCfg.token = token
	botCfg.groupID = groupID

	return &Bot{
		cfg: botCfg,
	}
}

func (b *Bot) SetTimeout(timeout int) {
	b.cfg.pollConfig.wait = timeout
}
