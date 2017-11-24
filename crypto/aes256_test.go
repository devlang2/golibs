package crypto

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := []byte("This is the key!") // 16 bytes
	data := []byte("Hello world~!")

	// Encrypt
	encrypted, err := EncAes256(key, data)
	if err != nil {
		t.Error(err.Error())
	}

	// Decrypt
	decrypted, err := DecAes256(key, encrypted)
	if err != nil {
		t.Error(err.Error())
	}

	// Compare
	if !bytes.Equal(data, decrypted) {
		fmt.Printf("plaintext: %x\n", data)
		fmt.Printf("decrypted: %x\n", decrypted)
		t.Error("Encrypted data is different from decrypted data")
	}
}
