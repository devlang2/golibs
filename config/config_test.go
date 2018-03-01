package config

import (
	"testing"
	"io/ioutil"
	"github.com/devplayg/golibs/crypto"
	"os"
)

func TestGetConfig(t *testing.T) {
	config := map[string]string{
		"what is your name?" : "devplayg",
		"what is your hobby?" : "programming",
		"what are you up to" : "programming",
	}

	configFile, err := ioutil.TempFile("", "config")
	defer os.Remove(configFile.Name()) // clean up
	if err != nil {
		t.Error(err)
	}

	encKey := []byte("this is the key!")
	err = crypto.SaveObjectToEncryptedFile(configFile.Name(), encKey, config)
	if err != nil {
		t.Error(err)
	}
	decryptedConfig := make(map[string]string)
	err = crypto.LoadEncryptedObjectFile(configFile.Name(), encKey, &decryptedConfig)
	if err != nil {
		t.Error(err)
	}

	for key, _ := range decryptedConfig {
		if _, ok := config[key]; !ok {
			t.Error("cannot find key")
		}
		if decryptedConfig[key] != config[key] {
			t.Error("cannot find key")
		}
	}
}