package caddydockerproxy

import (
	"testing"

	"github.com/lucaslorentz/caddy-docker-proxy/v2/config"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeAdminListen(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty",
			input:    "",
			expected: "",
		},
		{
			name:     "trim and add tcp prefix",
			input:    " 0.0.0.0:2019 ",
			expected: "tcp/0.0.0.0:2019",
		},
		{
			name:     "keep prefixed listen value",
			input:    "tcp/0.0.0.0:2019",
			expected: "tcp/0.0.0.0:2019",
		},
		{
			name:     "keep unix listen value",
			input:    "unix//run/caddy-admin.sock",
			expected: "unix//run/caddy-admin.sock",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, normalizeAdminListen(testCase.input))
		})
	}
}

func TestGetAdminListenPrefersConfiguredListen(t *testing.T) {
	options := &config.Options{
		AdminListen: "tcp/0.0.0.0:2019",
	}

	assert.Equal(t, "tcp/0.0.0.0:2019", getAdminListen(options))
}
