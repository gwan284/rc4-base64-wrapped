package rc4_base64_wrapped

import (
	"testing"
	"strings"
)

const TEST_KEY = "S3FL3Y8450"

func TestWebSafeBase64(t *testing.T) {
	t.Log("Coder should output web-safe base64 encoded string")

	c, _ := NewCoder(TEST_KEY)
	encoded := c.EncodeWrap("link-D5a1Z-user@yahoo.com")

	if strings.ContainsAny(encoded, "+,=") {
		t.Errorf("Expected URl-safe string, but it was %s instead.", encoded)
	}
}

func TestEncodingDecoding(t *testing.T) {
	t.Log("Encoding and decoding encoded string should give an original input sequence")

	str := "link-D5a1Z-user@yahoo.com"
	c, _ := NewCoder(TEST_KEY)

	encoded := c.EncodeWrap(str)
	decoded, _ := c.UnwrapDecode(encoded)

	if decoded != str {
		t.Errorf("Expected decoded string to be [%s], but it was [%s] instead.", str, decoded)
	}
}

func EncodeDecodeNTimes(s string, n int) string {
	decoded := ""

	c, _ := NewCoder(TEST_KEY)
	for i := 0; i < n; i++ {
		encoded := c.EncodeWrap(s)
		decoded, _ = c.UnwrapDecode(encoded)
	}
	return decoded
}

func EncodeNTimes(s string, n int) string {
	encoded := ""

	c, _ := NewCoder(TEST_KEY)
	for i := 0; i < n; i++ {
		encoded = c.EncodeWrap(s)
	}
	return encoded
}

func BenchmarkEncodingDecoding1Million(b *testing.B) {
	links := []string{
		"link-D5a1Z-small@gmail.com",
		"link-Adn54-medium.length@yahoo.com",
		"link-5MN6j-very_long_name.123@mailboxhostname.com",
	}

	for n := 0; n < b.N; n++ {
		EncodeDecodeNTimes(links[n%3], 1000000)
	}
}

func BenchmarkEncoding1Million(b *testing.B) {
	links := []string{
		"link-D5a1Z-small@gmail.com",
		"link-Adn54-medium.length@yahoo.com",
		"link-5MN6j-very_long_name.123@mailboxhostname.com",
	}

	for n := 0; n < b.N; n++ {
		EncodeNTimes(links[n%3], 1000000)
	}
}
