package cprotect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVaultReadWrite(t *testing.T) {
	content := "12345"
	vault := &PCVaultService{}
	err := vault.WriteToVault(test_product_code, content, false)
	assert.Nil(t, err)
	vault_content, err := vault.ReadVault(test_product_code, false)
	assert.Nil(t, err)
	assert.Equal(t, content, vault_content)
}
