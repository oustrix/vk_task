package bot

type botConfig struct {
	token string
}

func newConfig() *botConfig {
	return &botConfig{}
}
