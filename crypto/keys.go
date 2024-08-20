package crypto

// Keys lifted from https://github.com/phyber/avsg/blob/main/src/crypto.rs - thank you very much!
var (
	SAVEGAME_KEY = []byte{
		186, 173, 240, 13,
		0, 0, 0, 0,
		32, 48, 68, 194,
		19, 228, 31, 255,
	}

	SAVEGAME_IV = []byte{
		229, 255, 255, 255,
		229, 186, 7, 0,
		186, 173, 240, 13,
		255, 0, 255, 0,
	}
)
