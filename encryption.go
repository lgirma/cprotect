package cprotect

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var bytes = []byte{35, 12, 57, 24, 44, 34, 24, 74, 96, 35, 88, 85, 18, 96, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, password string) (string, error) {
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Decrypt(text, password string) (string, error) {
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	cipherText := decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
