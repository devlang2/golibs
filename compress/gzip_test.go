package compress

import (
	"bytes"
	"crypto"
	"errors"
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

func TestCompression(t *testing.T) {
	for str, _ := range Samples {
		if err := compress([]byte(str), GZIP); err != nil {
			t.Error(err)
		}
	}
}

func compress(data []byte, algorithm string) error {
	compressed, err := Compress(data, algorithm)
	if err != nil {
		return err
	}

	decompressed, err := Decompress(compressed, algorithm)
	if err != nil {
		return err
	}

	if !bytes.Equal(data, decompressed) {
		return errors.New("checksum error")
	}

	return nil
}
