package bot

type botConfig struct {
	token   string
	groupID string
}

func newConfig() *botConfig {
	return &botConfig{}
}
