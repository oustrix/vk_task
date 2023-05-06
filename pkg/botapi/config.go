package botapi

const (
	_LongPollWaitDefault = 25
)

type botConfig struct {
	token      string
	groupID    string
	pollConfig *longPollConfig
}

type longPollConfig struct {
	server string
	key    string
	ts     int
	wait   int
}

func newBotConfig() *botConfig {
	return &botConfig{}
}

func newLongPollConfig(server, key string, ts int) *longPollConfig {
	return &longPollConfig{
		server: server,
		key:    key,
		ts:     ts,
		wait:   _LongPollWaitDefault,
	}
}
