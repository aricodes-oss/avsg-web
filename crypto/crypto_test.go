package crypto

import (
	. "avsg/embeds"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecrypt(t *testing.T) {
	assert := assert.New(t)
	data, err := Decrypt(Samples.Encrypted)

	assert.Nil(err)
	assert.ElementsMatch(data, Samples.Decrypted)
	assert.Equal(data, Samples.Decrypted)
}

func TestEncrypt(t *testing.T) {
	assert := assert.New(t)
	data, err := Encrypt(Samples.Decrypted)

	assert.Nil(err)
	assert.ElementsMatch(data, Samples.Encrypted)
	assert.Equal(data, Samples.Encrypted)
}
