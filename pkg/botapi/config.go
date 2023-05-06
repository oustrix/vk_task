package botapi

const (
	_LongPollWaitDefault = 25
)

type longPollConfig struct {
	server string
	key    string
	ts     int
	wait   int
}

func newLongPollConfig(server, key string, ts int) *longPollConfig {
	return &longPollConfig{
		server: server,
		key:    key,
		ts:     ts,
		wait:   _LongPollWaitDefault,
	}
}
