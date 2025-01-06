package lib

import (
	"fmt"
	"strings"
)

type Formatter struct{}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) Format(offset int64, data []byte) string {
	hexPart := make([]string, len(data))
	asciiPart := make([]rune, len(data))

	for i, b := range data {
		hexPart[i] = fmt.Sprintf("%02x", b)
		if b >= 32 && b <= 126 {
			asciiPart[i] = rune(b)
		} else {
			asciiPart[i] = '.'
		}
	}

	return fmt.Sprintf("%08x  %-48s  |%s|\n", offset, strings.Join(hexPart, " "), string(asciiPart))
}
