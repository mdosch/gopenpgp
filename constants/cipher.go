package constants

// Cipher suite names.
const (
	ThreeDES  = "3des"
	TripleDES = "tripledes" // Both "3des" and "tripledes" refer to 3DES.
	CAST5     = "cast5"
	AES128    = "aes128"
	AES192    = "aes192"
	AES256    = "aes256"
)

const (
	SIGNATURE_OK          int = 0
	SIGNATURE_NOT_SIGNED  int = 1
	SIGNATURE_NO_VERIFIER int = 2
	SIGNATURE_FAILED      int = 3
	SIGNATURE_BAD_CONTEXT int = 4
)

// SecurityLevel constants
// The type is int8 for compatibility with
// gomobile.
const (
	StandardSecurity int8 = 0
	HighSecurity     int8 = 1
)
