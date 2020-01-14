package checksum

import (
	"crypto"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
)

func GetFileChecksum(path string, algorithm crypto.Hash) ([]byte, error) {
	if algorithm == crypto.MD5 {
		return getMd5Checksum(path)
	}
	if algorithm == crypto.SHA256 {
		return getSha256Checksum(path)
	}

	return nil, errors.New(fmt.Sprintf("not supported algorithm: %v", algorithm))
}

func getSha256Checksum(fp string) ([]byte, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func getMd5Checksum(fp string) ([]byte, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func GetStringChecksum(str string, algorithm crypto.Hash) ([]byte, error) {
	if algorithm == crypto.MD5 {
		h := md5.New()
		_, err := io.WriteString(h, str)
		if err != nil {
			return nil, err
		}
		return h.Sum(nil), nil
	}
	if algorithm == crypto.SHA256 {
		h := sha256.New()
		_, err := io.WriteString(h, str)
		if err != nil {
			return nil, err
		}
		return h.Sum(nil), nil
	}
	return nil, errors.New(fmt.Sprintf("not supported algorithm: %v", algorithm))
}
