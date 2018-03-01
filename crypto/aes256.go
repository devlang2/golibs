package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

func EncAes256(key, plaintext []byte) ([]byte, error) {
	plaintext = pad(plaintext)
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("Plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func DecAes256(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}

	if len(ciphertext) < aes.BlockSize {
		return []byte(""), errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	data := make([]byte, len(ciphertext)-aes.BlockSize)
	copy(data, ciphertext[aes.BlockSize:])

	if len(ciphertext[aes.BlockSize:])%aes.BlockSize != 0 {
		return []byte(""), errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, ciphertext[aes.BlockSize:])
	return unpad(data)
}

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}
	return src[:(length - unpadding)], nil
}

func SaveObjectToEncryptedFile(filePath string, key []byte, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	enc, err := EncAes256(key, b)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, enc, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadEncryptedObjectFile(filePath string, key []byte, v interface{}) error {
	enc, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	dec, err := DecAes256(key, enc)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dec, &v)
	return nil
}
