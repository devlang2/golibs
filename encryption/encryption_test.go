package encryption

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := []byte("This is the key!111") // 16 bytes
	data := []byte("Hello world~!")

	// Encrypt
	data_enc, err := Encrypt(key, data)
	if err != nil {
		t.Error(err.Error())
	}

	// Decrypt
	data_dec, err := Decrypt(key, data_enc)
	if err != nil {
		t.Error(err.Error())
	}

	// Compare
	if !bytes.Equal(data, data_dec) {
		fmt.Printf("plaintext: %x\n", data)
		fmt.Printf("decrypted: %x\n", data_dec)
		t.Error("Encrypted value is different from decrypted value")
	}
}
