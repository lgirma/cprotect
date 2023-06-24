package cprotect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	str := "abcdefg"
	rev := reverseString(str)
	assert.Equal(t, "gfedcba", rev)
}

func TestEncryption(t *testing.T) {
	password14 := "0a76b7c4275b4b"
	src := "cde7bc51"
	enc, err := Encrypt(src, password14)
	assert.Nil(t, err)
	assert.NotEmpty(t, enc)

	dec, err := Decrypt(enc, password14)
	assert.Nil(t, err)
	assert.NotEmpty(t, dec)
	assert.Equal(t, dec, src)	
}

func TestEncryptionWith32Password(t *testing.T) {
	scr2 := "fb65e876-3369-4e94-974b-222d36ff7655"
	password32 := "d4e22eb08dc048c1a9c7a3f44d5e0892"
	enc2, err := Encrypt(scr2, password32)
	assert.Nil(t, err)
	assert.NotEmpty(t, enc2)

	dec2, err := Decrypt(enc2, password32)
	assert.Nil(t, err)
	assert.NotEmpty(t, dec2)
	assert.Equal(t, dec2, scr2)
}

func TestNormalizePassword(t *testing.T) {
	normalized := normalize_password("123456789")
	assert.Len(t, normalized, 16)

	normalized = normalize_password("1234567891011121")
	assert.Len(t, normalized, 16)

	normalized = normalize_password("123456789101112145")
	assert.Len(t, normalized, 32)
}
