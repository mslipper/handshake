package primitives

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateBlind(t *testing.T) {
	tests := []struct {
		value    uint64
		nonceHex string
		blindHex string
	}{
		{
			410050,
			"fa3858ba5b513da9bc7276820d0d0e678da26ef47632e3d325caf2c20b877453",
			"7694e4a9c3598d64de5f834c5b0bce3d0bd31cdc9dbeb11571aeef4eda853108",
		},
	}
	for _, tt := range tests {
		nonceB, err := hex.DecodeString(tt.nonceHex)
		require.NoError(t, err)
		blindB, err := CreateBlind(tt.value, nonceB)
		require.Equal(t, tt.blindHex, hex.EncodeToString(blindB))
	}
}
