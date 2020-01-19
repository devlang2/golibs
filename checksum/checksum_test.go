package checksum

import (
	"crypto"
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

type checksumTest struct {
	str       string
	md5Sum    string
	sha256Sum string
}

var checksumTests = []checksumTest{
	{
		str:       "devplayg",
		md5Sum:    "751aaecf426dc3e74dfc3f28862b578a",
		sha256Sum: "4af8d36e7765aacf1b9eea6ff309f3fc446f192d277875a534b9504c1f4a8418",
	}, {
		str:       "one band one sound",
		md5Sum:    "07b2daa4115d93545d965ed1fd7ce351",
		sha256Sum: "e338d22b56f2d304312161306096a25a228c29ac54954cfab3aef73c194b6d8b",
	}, {
		str:       "What a wonderful world!",
		md5Sum:    "dc952f9ab1f161789f98ecca96a0db4b",
		sha256Sum: "1dee1fcf3a3ee25dce5fea9b320ecefe689377a0f27f262b3d5990441f7f6ed1",
	}, {
		str:       "안녕하세요, aloha, こんにちは你好",
		md5Sum:    "1b2b05cc80a7ff1023f7ad21a038fea3",
		sha256Sum: "87c73c7b60b4676a3fdc2fbed8888f41ed33b5a18d10998eaf2177b6c40e10e6",
	},
}

func TestChecksum(t *testing.T) {
	for _, tc := range checksumTests {
		if err := compareFileChecksum(tc.str, tc.md5Sum, tc.sha256Sum); err != nil {
			t.Error(err)
		}
	}
}

func compareFileChecksum(str, md5Sum, sha256Sum string) error {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	defer func() {
		os.Remove(f.Name())
	}()

	_, err = f.WriteString(str)
	if err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	var sum string
	sum, err = GetStringChecksumFromFile(f.Name(), crypto.MD5)
	if err != nil {
		return err
	}
	if sum != md5Sum {
		return errors.New("MD5 checksum failed")
	}

	sum, err = GetStringChecksumFromFile(f.Name(), crypto.SHA256)
	if err != nil {
		return err
	}
	if sum != sha256Sum {
		return errors.New("MD5 checksum failed")
	}
	return nil
}
