package botapi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewBot(t *testing.T) {
	r := require.New(t)

	type test struct {
		name     string
		token    string
		groupID  string
		expected *Bot
	}

	tests := []test{
		{name: "simple", token: "test_token", groupID: "test_group", expected: &Bot{token: "test_token", GroupID: "test_group"}},
		{name: "all empty", token: "", groupID: "", expected: &Bot{token: "", GroupID: ""}},
		{name: "token empty", token: "", groupID: "test_group", expected: &Bot{token: "", GroupID: "test_group"}},
		{name: "groupID empty", token: "test_token", groupID: "", expected: &Bot{token: "test_token", GroupID: ""}},
	}

	for _, tc := range tests {
		actual := NewBot(tc.token, tc.groupID)
		r.Equal(tc.expected.token, actual.token, tc.name)
		r.Equal(tc.expected.GroupID, actual.GroupID, tc.name)
	}
}
