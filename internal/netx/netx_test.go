package netx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBlockedLocalHostname(t *testing.T) {
	tests := []struct {
		name      string
		hostname  string
		allowlist []string
		want      bool
	}{
		{name: "loopback IPv4", hostname: "127.0.0.1", want: true},
		{name: "loopback IPv6", hostname: "::1", want: true},
		{name: "expanded loopback IPv6", hostname: "0:0:0:0:0:0:0:1", want: true},
		{name: "loopback IPv4 range", hostname: "127.0.0.95", want: true},
		{name: "this network", hostname: "0.0.0.0", want: true},
		{name: "private IPv4", hostname: "192.168.123.45", want: true},

		{name: "public IPv4", hostname: "165.232.140.255", want: false},
		{name: "lookup failure", hostname: "not a valid hostname", want: true},

		{name: "local IPv4 not allowlisted", hostname: "192.168.123.45", allowlist: []string{"10.0.0.17"}, want: true},
		{name: "allowlisted name", hostname: "gogs.local", allowlist: []string{"gogs.local"}, want: false},

		{name: "wildcard allowlist", hostname: "192.168.123.45", allowlist: []string{"*"}, want: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, IsBlockedLocalHostname(test.hostname, test.allowlist))
		})
	}
}
