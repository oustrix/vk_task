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
