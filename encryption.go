package cprotect

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"math"
)

const password_fill = "52624fb7e18e45a"

func normalize_password(password string) string {
	plen := len(password)
	factor := int(math.Ceil(float64(plen) / 16))
	remaining := int(math.Abs(float64(plen - 16*factor)))
	return password + password_fill[:remaining]
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func reverseString(s string) string {
	runes := []rune(s)
	size := len(runes)
	for i := 0; i < size/2; i++ {
		runes[size-i-1], runes[i] = runes[i], runes[size-i-1]
	}
	return string(runes)
}

func Encrypt(text, password_input string) (string, error) {
	password := normalize_password(password_input)
	iv := []byte(reverseString(password))
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	if len(iv) < block.BlockSize() {
		for i := 0; i < block.BlockSize()-len(iv); i++ {
			iv = append(iv, 0)
		}
	} else if len(iv) > block.BlockSize() {
		iv = iv[:block.BlockSize()]
	}
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

func Decrypt(text, password_input string) (string, error) {
	password := normalize_password(password_input)
	iv := []byte(reverseString(password))
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}
	cipherText := decode(text)
	if len(iv) < block.BlockSize() {
		for i := 0; i < block.BlockSize()-len(iv); i++ {
			iv = append(iv, 0)
		}
	} else if len(iv) > block.BlockSize() {
		iv = iv[:block.BlockSize()]
	}
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
