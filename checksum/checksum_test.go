package checksum

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileToSha256(t *testing.T) {
	cases := map[string]string{
		"devplayg":                "4AF8D36E7765AACF1B9EEA6FF309F3FC446F192D277875A534B9504C1F4A8418",
		"one band one sound":      "E338D22B56F2D304312161306096A25A228C29AC54954CFAB3AEF73C194B6D8B",
		"What a wonderful world!": "1DEE1FCF3A3EE25DCE5FEA9B320ECEFE689377A0F27F262B3D5990441F7F6ED1",
		"안녕하세요,aloha,こんにちは你好": "E99201CD3BE720A58E7038B208C423900B7D423587531764160AEB8484097471",
	}

	for k, v := range cases {
		if err := compareSha256(k, v); err != nil {
			t.Error(err)
		}
	}
}

func compareSha256(keyword, hexstr string) error {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	file.WriteString(keyword)
	file.Close()

	fileHash, err := FileToSha256(file.Name())
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(hexstr)
	if err != nil {
		return err
	}
	if bytes.Compare(b, fileHash) != 0 {

		return errors.New("not matched")
	}

	return nil
}
