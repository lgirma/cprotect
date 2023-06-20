package cprotect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	src := "cde7bc51"
	enc, err := Encrypt(src, test_password)
	assert.Nil(t, err)
	assert.NotEmpty(t, enc)

	dec, err := Decrypt(enc, test_password)
	assert.Nil(t, err)
	assert.NotEmpty(t, dec)
	assert.Equal(t, dec, src)
}

func TestNormalizePassword(t *testing.T) {
	normalized := normalize_password("123456789")
	assert.Len(t, normalized, 16)

	normalized = normalize_password("1234567891011121")
	assert.Len(t, normalized, 16)

	normalized = normalize_password("123456789101112145")
	assert.Len(t, normalized, 32)
}