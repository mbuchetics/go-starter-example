package util_test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-starter-example/internal/util"
)

func TestGenerateRandom(t *testing.T) {
	t.Parallel()

	res, err := util.GenerateRandomBytes(13)
	require.NoError(t, err)
	assert.Len(t, res, 13)

	randString, err := util.GenerateRandomBase64String(17)
	require.NoError(t, err)
	res, err = base64.StdEncoding.DecodeString(randString)
	require.NoError(t, err)
	assert.Len(t, res, 17)

	randString, err = util.GenerateRandomHexString(19)
	require.NoError(t, err)
	res, err = hex.DecodeString(randString)
	require.NoError(t, err)
	assert.Len(t, res, 19)
}
