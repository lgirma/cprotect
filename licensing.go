package cprotect

import "errors"

func IsInstalled(product string, password string, forAllUsers bool, fingerprintService FingerprintService, vault VaultService) (bool, error) {
	reqCode, err := GetRequestCode(product, fingerprintService)
	if err != nil {
		return false, err
	}
	vaultContent, err := vault.ReadVault(product, forAllUsers)
	if len(vaultContent) == 0 {
		return false, err
	}
	dec, err := Decrypt(vaultContent, password)
	if err != nil {
		return false, errors.New(ErrorDecryptionFailure)
	}
	return dec == reqCode, nil
}

func IsActivationCodeValid(password string, requestCode string, activationCode string) bool {
	dec, _ := Decrypt(activationCode, password)
	return dec == requestCode
}

func Install(product string, activationCode string, forAllUsers bool, vault VaultService) error {
	return vault.WriteToVault(product, activationCode, forAllUsers)
}

func Uninstall(product string, forAllUsers bool, vault VaultService) error {
	return vault.WriteToVault(product, "", forAllUsers)
}