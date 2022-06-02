package cprotect

import (
	"errors"
	"os"
	"path"
)

func getVaultFilePaths(product string) []string {
	result := make([]string, 0)
	configDir, err := os.UserConfigDir()
	if err == nil {
		result = append(result, path.Join(configDir, product, "vault.key"))
	}

	return result
}

func readVaultFile(product string) (string, error) {
	paths := getVaultFilePaths(product)
	vaultFilePath := ""
	for _, filePath := range paths {
		_, err := os.Stat(filePath)
		if err == nil {
			vaultFilePath = filePath
			break
		}
	}
	if len(vaultFilePath) == 0 {
		return "", nil
	}
	content, err := os.ReadFile(vaultFilePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func IsInstalled(product string, password string) (bool, error) {
	reqCode, err := GetRequestCode(product)
	if err != nil {
		return false, err
	}
	vaultContent, err := readVaultFile(product)
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

func Install(product string, activationCode string) error {
	vaultFiles := getVaultFilePaths(product)
	if len(vaultFiles) == 0 {
		return errors.New(ErrorSuitableVaultDirNotFound)
	}
	vaultFile := vaultFiles[0]
	dir := path.Dir(vaultFile)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return errors.New(ErrorSuitableVaultDirNotFound)
	}
	f, err := os.OpenFile(vaultFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New(ErrorVaultFileWriteFailure)
	}
	_, err = f.WriteString(activationCode)
	if err != nil {
		return errors.New(ErrorVaultFileWriteFailure)
	}
	defer f.Close()
	return nil
}