package crypto

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := []byte("AES256ENCRYPTION") // 16 or 32 bytes
	data := []byte("Hello gophers")

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

	// Save object to encrypted file
	tempDir := tempMkdir(t)
	filepath := tempMkFile(t, tempDir)
	defer os.RemoveAll(tempDir)

	// Save object to encrypted file
	type Obj struct {
		Name string
		Age  int
	}
	obj := Obj{"Golang", 9}
	err = SaveObjectToEncryptedFile(filepath, key, obj)
	if err != nil {
		t.Error(err.Error())
	}

	// Load encrypted file and
	objDecrypted := Obj{}
	err = LoadEncryptedObjectFile(filepath, key, &objDecrypted)
	if err != nil {
		t.Error(err.Error())
	}

	// Compare
	if !reflect.DeepEqual(obj, objDecrypted) {
		t.Error("Failed to decrypting file")
	}
}

func tempMkdir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "devplayg")
	if err != nil {
		t.Fatalf("failed to create test directory: %s", err)
	}
	return dir
}

func tempMkFile(t *testing.T, dir string) string {
	f, err := ioutil.TempFile(dir, "devplayg")
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer f.Close()
	return f.Name()
}
