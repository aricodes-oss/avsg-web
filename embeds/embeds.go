// Package embeds provides a single location to embed files used in testing.
package embeds

import (
	_ "embed"
)

//go:embed sample_encrypted.sav
var sampleEncrypted []byte

//go:embed sample_decrypted.xml
var sampleDecrypted []byte

var Samples = struct {
	Encrypted []byte
	Decrypted []byte
}{sampleEncrypted, sampleDecrypted}
