package rc4_base64_wrapped

import (
	"encoding/base64"
	"crypto/rc4"
	"errors"
)

type webSafeCoder struct {
	rc4key []byte
	b64 *base64.Encoding
}

const SprinkleChar = "/"

// Creates new coder using RC4-key provided and web-safe base64 encoder.
func NewCoder(key string) (*webSafeCoder, error) {
	var coder webSafeCoder

	coder.rc4key = []byte(key)
	if len(coder.rc4key) < 1 || len(coder.rc4key) > 256 {
		return nil, errors.New("Invalid RC4 crypto key size.")
	}

	coder.b64 = base64.URLEncoding.WithPadding(base64.NoPadding)

	return &coder, nil
}

//Encodes string using RC4 crypto algorithm
// and wraps resulting bytes using web-safe base64 encoding
func (c webSafeCoder) EncodeWrap(s string) string {
	data := []byte(s)
	encrypted := make([]byte, len(data))

	cipher, _ := rc4.NewCipher(c.rc4key)
	cipher.XORKeyStream(encrypted, data)
	cipher.Reset()

	encoded := c.b64.EncodeToString(encrypted)

	return sprinkle(encoded, SprinkleChar)
}

//Unwraps string using web-safe base64 encoding
// and decrypts result using RC4 cipher
func (c webSafeCoder) UnwrapDecode(s string) (string, error) {
	unsprikled := unsprinkle(s, SprinkleChar)

	decoded, err := c.b64.DecodeString(unsprikled)
	if err != nil {
		return s, err
	}

	decrypted := make([]byte, len(decoded))

	cipher, _ := rc4.NewCipher(c.rc4key)
	cipher.XORKeyStream(decrypted, decoded)
	cipher.Reset()

	return string(decrypted), err
}
