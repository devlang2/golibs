package secureconfig

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/devplayg/golibs/crypto"
)

func GetConfig(configPath string, key []byte) (map[string]string, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New("configuration file not found: "+configPath)
	} else {
		config := make(map[string]string)
		err := crypto.LoadEncryptedObjectFile(configPath, key, &config)
		return config, err
	}
}

func SetConfig(configPath, keys string, key []byte) error {
	config, _ := GetConfig(configPath, key)
	if config == nil {
		config = make(map[string]string)
	}

	fmt.Println("Setting configuration (for deleting: '' or \"\")")
	if len(keys) > 0 {
		arr := strings.Split(keys, ",")
		for _, k := range arr {
			readInput(strings.TrimSpace(k), config)
		}
	}
	err := crypto.SaveObjectToEncryptedFile(configPath, key, config)
	return err
}

func readInput(key string, config map[string]string) {
	if val, ok := config[key]; ok && len(val) > 0 {
		fmt.Printf("# %-16s = (%s) ", key, val)
	} else {
		fmt.Printf("# %-16s = ", key)
	}

	reader := bufio.NewReader(os.Stdin)
	newVal, _ := reader.ReadString('\n')
	newVal = strings.TrimSpace(newVal)
	if len(newVal) > 0 {
		if newVal == "\"\"" || newVal == "''" {
			config[key] = ""
		} else {
			config[key] = newVal
		}
	}
}
