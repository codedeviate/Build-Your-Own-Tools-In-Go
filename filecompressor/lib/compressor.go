// compressor.go
package lib

import (
	"compress/gzip"
	"io"
	"os"
)

// CompressFile compresses the given file using Gzip.
func CompressFile(source, target string) error {
    srcFile, err := os.Open(source)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    tgtFile, err := os.Create(target)
    if err != nil {
        return err
    }
    defer tgtFile.Close()

    writer := gzip.NewWriter(tgtFile)
    defer writer.Close()

    _, err = io.Copy(writer, srcFile)
    return err
}

// DecompressFile decompresses the given Gzip file.
func DecompressFile(source, target string) error {
    srcFile, err := os.Open(source)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    reader, err := gzip.NewReader(srcFile)
    if err != nil {
        return err
    }
    defer reader.Close()

    tgtFile, err := os.Create(target)
    if err != nil {
        return err
    }
    defer tgtFile.Close()

    _, err = io.Copy(tgtFile, reader)
    return err
}
