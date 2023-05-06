package bot

type botConfig struct {
	token   string
	groupID string
}

func newBotConfig() *botConfig {
	return &botConfig{}
}
