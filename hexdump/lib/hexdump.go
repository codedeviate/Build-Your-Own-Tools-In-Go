package lib

import (
	"io"
)

// Dump reads from the input reader and writes a formatted hexdump to the output writer.
func Dump(input io.Reader, output io.Writer) error {
	reader := NewReader(input)
	formatter := NewFormatter()

	var baseCount int64 = 0

	for {
		offset, data, err := reader.ReadChunk()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		line := formatter.Format(baseCount, data)
		baseCount += offset
		_, err = output.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return nil
}
