// encodedecode.go
package lib

import (
	"encoding/base64"
	"net/url"
)

// Base64Encode encodes a string to Base64.
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64Decode decodes a Base64 string.
func Base64Decode(data string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// URLEncode encodes a string to URL encoding.
func URLEncode(data string) string {
	return url.QueryEscape(data)
}

// URLDecode decodes a URL encoded string.
func URLDecode(data string) (string, error) {
	decoded, err := url.QueryUnescape(data)
	if err != nil {
		return "", err
	}
	return decoded, nil
}
