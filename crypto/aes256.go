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

func pkcsPadding(data []byte, blockSize int) []byte {
	paddingLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	return append(data, padding...)
}

func pkcsUnpadding(data []byte) []byte {
	length := len(data)
	paddingLen := int(data[length-1])
	return data[:(length - paddingLen)]
}

//func AES256(data []byte, key [32]byte) ([]byte, error) {
//
//}

//func aes256(data []byte, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create cipher block for AES256")
//	}
//
//	blockSize := 32
//	if block.BlockSize() != blockSize {
//		return nil, fmt.Errorf("block size of AES256 must be %d", blockSize)
//	}
//
//	return aesCBCEncrypt(data, block)
//}
//
//func aes128(data []byte, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create cipher block for AES256")
//	}
//
//	blockSize := 32
//	if block.BlockSize() != blockSize {
//		return nil, fmt.Errorf("block size of AES256 must be %d", blockSize)
//	}
//
//	return aesCBCEncrypt(data, block)
//}

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

func aesCBCEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Pad data
	data = pkcsPadding(data, block.BlockSize())
	encData := make([]byte, block.BlockSize()+len(data))

	// Generate IV
	iv := encData[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Crypt blocks
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

	//data := make([]byte, len(ciphertext)-aes.BlockSize)
	//copy(data, ciphertext[aes.BlockSize:])
	//

	//
	//mode := cipher.NewCBCDecrypter(block, iv)
	//mode.CryptBlocks(data, ciphertext[aes.BlockSize:])
	//return unpad(data)

	//
	//	mode := cipher.NewCBCDecrypter(block, iv)
	//
	//	// CryptBlocks can work in-place if the two arguments are the same.
	//	mode.CryptBlocks(encryptData, encryptData)
	//	/ / Unfill
	//	encryptData = PKCS7UnPadding(encryptData)
	//	return encryptData,nil
}

//func EncAes256(key, plaintext []byte) ([]byte, error) {
//	plaintext = pad(plaintext)
//	if len(plaintext)%aes.BlockSize != 0 {
//		return nil, errors.New("Plaintext is not a multiple of the block size")
//	}
//
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//
//	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
//	iv := ciphertext[:aes.BlockSize]
//	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//		panic(err)
//	}
//
//	mode := cipher.NewCBCEncrypter(block, iv)
//	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
//	return ciphertext, nil
//}

//func AesCBCEncrypt(rawData,key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}
//	blockSize := block.BlockSize()
//	rawData = PKCS7Padding(rawData, blockSize)
//	cipherText := make([]byte,blockSize+len(rawData))
//	iv := cipherText[:blockSize]
//	if _, err := io.ReadFull(rand.Reader,iv); err != nil {
//		panic(err)
//	}
//
//	//block size and initial vector size must be the same
//	mode := cipher.NewCBCEncrypter(block,iv)
//	mode.CryptBlocks(cipherText[blockSize:],rawData)
//
//	return cipherText, nil
//}
//
//func AesCBCDncrypt(encryptData, key []byte) ([]byte,error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}
//
//	blockSize := block.BlockSize()
//
//	if len(encryptData) < blockSize {
//		panic("ciphertext too short")
//	}
//	iv := encryptData[:blockSize]
//	encryptData = encryptData[blockSize:]
//
//	// CBC mode always works in whole blocks.
//	if len(encryptData)%blockSize != 0 {
//		panic("ciphertext is not a multiple of the block size")
//	}
//
//	mode := cipher.NewCBCDecrypter(block, iv)
//
//	// CryptBlocks can work in-place if the two arguments are the same.
//	mode.CryptBlocks(encryptData, encryptData)
//	/ / Unfill
//	encryptData = PKCS7UnPadding(encryptData)
//	return encryptData,nil
//}
//
