package rc4_base64_wrapped

import (
	"testing"
	"strings"
)

func TestWebSafebase64(t *testing.T) {
	t.Log("Coder should output web-safe base64 encoded string")

	c, _ := NewCoder("S3FL3Y8450")
	encoded := c.EncodeWrap("link-D5a1Z-user@yahoo.com")

	if strings.ContainsAny(encoded, "+,/=") {
		t.Errorf("Expected URl-safe string, but it was %s instead.", encoded)
	}
}

func TestEncodingDecoding(t *testing.T) {
	t.Log("Encoding and decoding encoded string should give an original input sequence")

	str := "link-D5a1Z-user@yahoo.com"
	c, _ := NewCoder("S3FL3Y8450")

	encoded := c.EncodeWrap(str)
	decoded, _ := c.UnwrapDecode(encoded)

	if decoded != str {
		t.Errorf("Expected decoded string to be [%s], but it was [%s] instead.", str, decoded)
	}
}