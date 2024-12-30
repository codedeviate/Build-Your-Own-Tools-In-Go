// hash.go
package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

// HashData hashes the input data using the specified hash algorithm.
func HashData(data string, algorithm string) (string, error) {
    var hasher hash.Hash

    switch algorithm {
    case "md5":
        hasher = md5.New()
    case "sha1":
        hasher = sha1.New()
    case "sha256":
        hasher = sha256.New()
    case "sha512":
        hasher = sha512.New()
    default:
        return "", fmt.Errorf("unsupported hash algorithm: %s", algorithm)
    }

    _, err := hasher.Write([]byte(data))
    if err != nil {
        return "", err
    }

    return hex.EncodeToString(hasher.Sum(nil)), nil
}

// HashFile hashes the content of the specified file using the specified hash algorithm.
func HashFile(filePath string, algorithm string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    var hasher hash.Hash

    switch algorithm {
    case "md5":
        hasher = md5.New()
    case "sha1":
        hasher = sha1.New()
    case "sha256":
        hasher = sha256.New()
    case "sha512":
        hasher = sha512.New()
    default:
        return "", fmt.Errorf("unsupported hash algorithm: %s", algorithm)
    }

    if _, err := io.Copy(hasher, file); err != nil {
        return "", err
    }

    return hex.EncodeToString(hasher.Sum(nil)), nil
}
