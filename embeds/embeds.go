// Package embeds provides a single location to embed files used in testing.
package embeds

import (
	"embed"
)

//go:embed sample_encrypted.sav
var sampleEncrypted []byte

//go:embed sample_decrypted.xml
var sampleDecrypted []byte

//go:embed dist
var Frontend embed.FS

var Samples = struct {
	Encrypted []byte
	Decrypted []byte
}{sampleEncrypted, sampleDecrypted}
