package crypto

import (
	_ "embed"

	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample_decrypted.xml
var sampleDecrypted []byte

//go:embed sample_encrypted.sav
var sampleEncrypted []byte

func TestDecrypt(t *testing.T) {
	assert := assert.New(t)
	data, err := Decrypt(sampleEncrypted)

	assert.Nil(err)
	assert.ElementsMatch(data, sampleDecrypted)
	assert.Equal(data, sampleDecrypted)
}

func TestEncrypt(t *testing.T) {
	assert := assert.New(t)
	data, err := Encrypt(sampleDecrypted)

	assert.Nil(err)
	assert.ElementsMatch(data, sampleEncrypted)
	assert.Equal(data, sampleEncrypted)
}
