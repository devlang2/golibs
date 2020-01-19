package checksum

import (
	"crypto"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
)

func GetStringChecksumFromFile(path string, hashType crypto.Hash) (string, error) {
	b, err := GetBytesChecksumFromFile(path, hashType)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func GetBytesChecksumFromFile(path string, hashType crypto.Hash) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var h hash.Hash
	if hashType == crypto.MD5 {
		h = md5.New()
	} else if hashType == crypto.SHA256 {
		h = sha256.New()
	} else {
		return nil, errors.New(fmt.Sprintf("not supported Hash: %d", hashType))
	}

	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	sum := h.Sum(nil)
	return sum[:], nil
}
