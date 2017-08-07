package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	//	"github.com/davecgh/go-spew/spew"
)

//var (
//	key = []byte("c43ac86d84469030f28c0a9656b1c533")
//	iv  = []byte("2981eeca66b5c3cd")
//)

//func main() {
//	//	spew.Dump()
//	key := []byte("c43ac86d84469030f28c0a9656b1c533")
//	iv := []byte("2981eeca66b5c3cd")

//	//	var str = "Hello world!  "
//	//	data := []byte(str)
//	//	data_enc, _ := Encrypt(key, data)
//	//	spew.Dump(data_enc)
//	//	data_dec, _ := Decrypt(key, data_enc)
//	//	spew.Dump(data_dec)

//	data_enc, _ := hex.DecodeString("5f3d4526d15a37cf8243103b6004b3a13ff8abe735ecc788d4879f3bef34a92ce446cb97aed9350704351b27dfb7e851991ad101b0be39154165c61856be2f178513d057024eb8b628dfca77607742d68206c20667b6a54fb467bdbbd2df71ab1e4430bf4ad279db3d08332c55d12f05e1e996a46d11d9c753f845eb87b1c1189f0b3af3057c9dd657fbde1ac637cf62")
//	spew.Dump(data_enc)
//	src := append(iv, data_enc...)
//	data_dec, _ := Decrypt(key, src)
//	spew.Dump(data_dec)
//	arr := bytes.Split(data_dec, []byte("|"))
//	spew.Dump(arr)

//	//	data := getData()

//	//	decrypt(data)
//	//	str := decrypt(data)
//	//	spew.Dump(str)
//}

func Encrypt(key, plaintext []byte) ([]byte, error) {
	plaintext = pad(plaintext)
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
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

func Decrypt(key, ciphertext []byte) ([]byte, error) {
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
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("Unpad error. This could happen when incorrect encryption key is used")
	}
	return src[:(length - unpadding)], nil
}
