package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func _aesCBCEncrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Block size and IV size must be the same
	if len(iv) != block.BlockSize() {
		return nil, errors.New("block size and IV size must be the same")
	}

	// Pad data
	data = pkcsPadding(data, block.BlockSize())
	encData := make([]byte, len(data))
	encData = append(iv, encData...)

	// Crypt blocks
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(encData[block.BlockSize():], data)
	return encData, nil
}

func EncAes256Old(key, iv, plaintext []byte) ([]byte, error) {
	plaintext = pad(plaintext)
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("Plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	// ciphertext = append(iv, ciphertext)
	encData := make([]byte, len(plaintext))
	ciphertext := append(iv, encData...)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// func p(str string) {
// 	fmt.Println("######################### "+str+" #########################")
// }
func aesCBCEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	p("block-size")
	fmt.Printf("block-size=%d, key-size=%d\n", block.BlockSize(), len(key))

	// Pad data
	data = pkcsPadding(data, block.BlockSize())
	encData := make([]byte, block.BlockSize()+len(data))

	// Generate IV
	iv := encData[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Encrypt blocks
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(encData[block.BlockSize():], data)
	return encData, nil
}

func aesCBCDncrypt(encData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(encData) < block.BlockSize() {
		return nil, errors.New("ciphertext too short")
	}

	iv := encData[:block.BlockSize()]

	encData = encData[block.BlockSize():]
	if len(encData)%block.BlockSize() != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encData, encData)

	return pkcsUnpadding(encData), nil

	// data := make([]byte, len(ciphertext)-aes.BlockSize)
	// copy(data, ciphertext[aes.BlockSize:])
	//

	//
	// mode := cipher.NewCBCDecrypter(block, iv)
	// mode.CryptBlocks(data, ciphertext[aes.BlockSize:])
	// return unpad(data)

	//
	//	mode := cipher.NewCBCDecrypter(block, iv)
	//
	//	// CryptBlocks can work in-place if the two arguments are the same.
	//	mode.CryptBlocks(encryptData, encryptData)
	//	/ / Unfill
	//	encryptData = PKCS7UnPadding(encryptData)
	//	return encryptData,nil
}
