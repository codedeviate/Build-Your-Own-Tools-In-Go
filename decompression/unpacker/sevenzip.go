// unpacker/sevenzip.go
package unpacker

import (
	"fmt"

	"github.com/mholt/archiver/v3"
)

// Un7z extracts a .7z file to the specified destination folder.
func Un7z(src string, dest string) error {
	// Use the archiver package to extract the .7z file
	err := archiver.Unarchive(src, dest)
	if err != nil {
		return fmt.Errorf("failed to unpack .7z file: %v", err)
	}
	return nil
}
