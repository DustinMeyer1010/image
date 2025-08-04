package png

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/stretchr/testify/require"
)

// Normal png pull from the internet
func TestValidPNG(t *testing.T) {

	png, err := CreatePNG("../testimages/valid copy.png")

	require.NoError(t, err)
	require.NotNil(t, png)
	assert.Equal(t, png.Signature, PNG_SINGATURE)
}

// invalidSignture has one byte of signature replace with randon byte
func TestInvalid(t *testing.T) {
	png, err := CreatePNG("../testimages/invalidSignature.png")

	require.Error(t, err)
	require.Nil(t, png)
}

func TestMissingIEND(t *testing.T) {
	png, err := CreatePNG("../testimages/invalidMissingIEND.png")

	require.Error(t, err)
	require.Nil(t, png)
}
