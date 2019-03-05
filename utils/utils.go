package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// MD5CreateStrings ...
func MD5CreateStrings(plaintext string) string {
	m := md5.New()

	plainbyte := []byte(plaintext)
	m.Write(plainbyte)

	cipherbyte := m.Sum(nil)
	ciphertext := hex.EncodeToString(cipherbyte)
	return ciphertext
}

// PKCS5Padding ...
func PKCS5Padding(cipherbyte []byte, blockSize int) []byte {
	padding := blockSize - len(cipherbyte)%blockSize
	paddingbyte := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherbyte, paddingbyte...)
}

// PKCS5UnPadding ...
func PKCS5UnPadding(plainbyte []byte) []byte {
	length := len(plainbyte)
	unpadding := int(plainbyte[length-1])
	return plainbyte[:(length - unpadding)]
}

func AesEncrypt(plainbyte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	plainbyte = PKCS5Padding(plainbyte, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	cipherbyte := make([]byte, len(plainbyte))
	blockMode.CryptBlocks(cipherbyte, plainbyte)

	return base64.StdEncoding.EncodeToString(cipherbyte), nil
}

func AesDecrypt(ciphertext string, key []byte) (string, error) {
	cipherbyte, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cipherbyte))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
