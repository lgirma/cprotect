package cprotect

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func reverseString(s string) string {
	runes := []rune(s)
	size := len(runes)
	for i := 0; i < size/2; i++ {
		runes[size-i-1], runes[i] = runes[i],  runes[size-i-1]
	}
	return string(runes)
}


func Encrypt(text, password string) (string, error) {
	iv := []byte(reverseString(password))
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func Decrypt(text, password string) (string, error) {
	iv := []byte(reverseString(password))
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	cipherText := decode(text)
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
