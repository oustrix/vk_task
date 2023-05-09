package botapi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLongPollConfig(t *testing.T) {
	r := require.New(t)

	type test struct {
		name     string
		server   string
		key      string
		ts       int
		expected *longPollConfig
	}

	tests := []test{
		{name: "simple", server: "server123", key: "key123", ts: 123, expected: &longPollConfig{server: "server123", key: "key123", ts: 123, Wait: _LongPollWaitDefault}},
		{name: "server empty", server: "", key: "key123", ts: 123, expected: &longPollConfig{server: "", key: "key123", ts: 123, Wait: _LongPollWaitDefault}},
		{name: "key empty", server: "server123", key: "", ts: 123, expected: &longPollConfig{server: "server123", key: "", ts: 123, Wait: _LongPollWaitDefault}},
		{name: "server and key are empty", server: "", key: "", ts: 123, expected: &longPollConfig{server: "", key: "", ts: 123, Wait: _LongPollWaitDefault}},
	}

	for _, tc := range tests {
		actual := newLongPollConfig(tc.server, tc.key, tc.ts)
		r.Equal(tc.expected, actual, tc.name)
	}
}
