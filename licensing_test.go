package cprotect

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummy_fingerprint_service struct {
}

func (service *dummy_fingerprint_service) GetMachineId(elevatedPrivilege bool) (string, error) {
	return "5f4a966c-b022-490c-92e5-86dfae6e671b", nil
}

var dummy_fp_service = &dummy_fingerprint_service{}

type dummy_vault_service struct {
	vault_content string
}

func (service *dummy_vault_service) ReadVault(product string, forAllUsers bool) (string, error) {
	return service.vault_content, nil
}

func (service *dummy_vault_service) WriteToVault(product string, content string, forAllUsers bool) error {
	service.vault_content = content
	return nil
}

var dummy_vault = &dummy_vault_service{}

const test_password = "0a76b7c4275b4b"
const test_product_code = "vault_test"

func TestGetRequestCode(t *testing.T) {
	reqCode, err := GetRequestCode(test_product_code, dummy_fp_service)
	assert.Nil(t, err)
	assert.NotNil(t, reqCode)
}

func TestInstall(t *testing.T) {
	activation_code := "activation_code"
	err := Install(test_product_code, activation_code, false, dummy_vault)
	assert.Nil(t, err)

	vault_content, err := dummy_vault.ReadVault(test_product_code, false)
	assert.Nil(t, err)
	assert.Equal(t, vault_content, activation_code)
}

func TestIsInstalled(t *testing.T) {
	dummy_vault.WriteToVault(test_product_code, "", false)
	installed, err := IsInstalled(test_product_code, test_password, false, dummy_fp_service, dummy_vault)
	assert.Nil(t, err)
	assert.False(t, installed)

	req_code, _ := GetRequestCode(test_product_code, dummy_fp_service)
	activation_code, _ := Encrypt(req_code, test_password)
	Install(test_product_code, activation_code, false, dummy_vault)

	installed, err = IsInstalled(test_product_code, test_password, false, dummy_fp_service, dummy_vault)
	assert.Nil(t, err)
	assert.True(t, installed)
}

func TestUninstall(t *testing.T) {
	req_code, _ := GetRequestCode(test_product_code, dummy_fp_service)
	activation_code, _ := Encrypt(req_code, test_password)
	Install(test_product_code, activation_code, false, dummy_vault)
	Uninstall(test_product_code, false, dummy_vault)

	vault_content, _ := dummy_vault.ReadVault(test_product_code, false)
	assert.Empty(t, vault_content)
}

func TestIsActivationCodeValid(t *testing.T) {
	req_code, _ := GetRequestCode(test_product_code, dummy_fp_service)
	activation_code, _ := Encrypt(req_code, test_password)
	invalid_code := strings.Repeat("a", len(activation_code))
	valid := IsActivationCodeValid(test_password, req_code, activation_code)
	invalid := IsActivationCodeValid(test_password, req_code, invalid_code)

	assert.True(t, valid)
	assert.False(t, invalid)
}
