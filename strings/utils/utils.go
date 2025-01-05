// utils/utils.go
package utils

import "unicode"

// isPrintable checks if a byte is a printable ASCII character.
func IsPrintable(c byte) bool {
	return unicode.IsPrint(rune(c)) && !unicode.IsControl(rune(c))
}
