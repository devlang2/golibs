package checksum

import (
	"crypto"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

var (
	Algorithms = [2]crypto.Hash{crypto.MD5, crypto.SHA256}
	Samples    = map[string][2]string{
		"devplayg":                {"751aaecf426dc3e74dfc3f28862b578a", "4af8d36e7765aacf1b9eea6ff309f3fc446f192d277875a534b9504c1f4a8418"},
		"one band one sound":      {"07b2daa4115d93545d965ed1fd7ce351", "e338d22b56f2d304312161306096a25a228c29ac54954cfab3aef73c194b6d8b"},
		"What a wonderful world!": {"dc952f9ab1f161789f98ecca96a0db4b", "1dee1fcf3a3ee25dce5fea9b320ecefe689377a0f27f262b3d5990441f7f6ed1"},
		"안녕하세요, aloha, こんにちは你好": {"1b2b05cc80a7ff1023f7ad21a038fea3", "87c73c7b60b4676a3fdc2fbed8888f41ed33b5a18d10998eaf2177b6c40e10e6"},
	}
)

func TestChecksum(t *testing.T) {
	for i, algorithm := range Algorithms {
		for str, checksums := range Samples {
			if err := compareStringChecksum(str, checksums[i], algorithm); err != nil {
				t.Error(err)
			}

			if err := compareFileChecksum(str, checksums[i], algorithm); err != nil {
				t.Error(err)
			}
		}
	}
}

func compareStringChecksum(str, expected string, algo crypto.Hash) error {
	checksum, err := GetStringChecksum(str, algo)
	if err != nil {
		return err
	}

	checksumStr := hex.EncodeToString(checksum)
	if checksumStr != expected {
		return errors.New("checksum error")
	}
	return nil
}

func compareFileChecksum(str, expected string, algo crypto.Hash) error {
	f, err := createTempFile(str)
	if err != nil {
		return err
	}
	defer func() {
		if err := os.Remove(f.Name()); err != nil {
			panic(err)
		}
	}()
	checksum, err := GetFileChecksum(f.Name(), algo)
	if err != nil {
		return err
	}
	checksumStr := hex.EncodeToString(checksum)
	if checksumStr != expected {
		return errors.New("checksum error")
	}

	return nil
}

func createTempFile(text string) (*os.File, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	if _, err := f.WriteString(text); err != nil {
		return nil, err
	}

	return f, nil

}
