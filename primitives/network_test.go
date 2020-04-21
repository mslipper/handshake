package primitives

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetwork_RPCPort(t *testing.T) {
	assert.Equal(t, 12037, NetworkMainnet.RPCPort())
	assert.Equal(t, 13037, NetworkTestnet.RPCPort())
	assert.Equal(t, 14037, NetworkRegtest.RPCPort())
	assert.Equal(t, 15037, NetworkSimnet.RPCPort())
	assert.Panics(t, func() {
		Network("foobar").RPCPort()
	})
}
