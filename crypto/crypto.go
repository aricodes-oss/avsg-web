package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

var block cipher.Block

func init() {
	b, err := aes.NewCipher(SAVEGAME_KEY)
	if err != nil {
		panic(err)
	}

	block = b
}

// Encrypt takes an unencrypted data buffer and returns the encrypted version
func Encrypt(data []byte) ([]byte, error) {
	// Add PKCS7 padding
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	data = append(data, padtext...)

	if len(SAVEGAME_IV) != aes.BlockSize {
		return nil, errors.New("invalid IV size")
	}

	cbc := cipher.NewCBCEncrypter(block, SAVEGAME_IV)
	encrypted := make([]byte, len(data))
	cbc.CryptBlocks(encrypted, data)

	return encrypted, nil
}

// Decrypt takes an encrypted data buffer and returns the unencrypted version
func Decrypt(data []byte) ([]byte, error) {
	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	if len(SAVEGAME_IV) != aes.BlockSize {
		return nil, errors.New("invalid IV size")
	}

	cbc := cipher.NewCBCDecrypter(block, SAVEGAME_IV)
	decrypted := make([]byte, len(data))
	cbc.CryptBlocks(decrypted, data)

	// Remove PKCS7 padding
	padding := int(decrypted[len(decrypted)-1])
	if padding > aes.BlockSize || padding > len(decrypted) {
		return nil, errors.New("invalid padding")
	}
	decrypted = decrypted[:len(decrypted)-padding]

	return decrypted, nil
}
