package cprotect

import (
	"errors"
	"os"
	"path/filepath"
)

type VaultService interface {
	ReadVault(product string, forAllUsers bool) (string, error)
	WriteToVault(product string, content string, forAllUsers bool) error
}

func GetVaultService() VaultService {
	return &PCVaultService{}
}

type PCVaultService struct {

}

const vault_file_name = "vault.key"

func (service *PCVaultService) GetVaultFilePaths(product string, forAllUsers bool) []string {
	result := make([]string, 0)
	configDir, err := os.UserConfigDir()
	if err == nil {
		result = append(result, filepath.Join(configDir, product, vault_file_name))
	}

	return result
}

func (service *PCVaultService) ReadVault(product string, forAllUsers bool) (string, error) {
	paths := service.GetVaultFilePaths(product, forAllUsers)
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

func (service *PCVaultService) WriteToVault(product string, content string, forAllUsers bool) error {
	vaultFiles := service.GetVaultFilePaths(product, true)
	if len(vaultFiles) == 0 {
		return errors.New(ErrorSuitableVaultDirNotFound)
	}
	vaultFile := vaultFiles[0]
	dir := filepath.Dir(vaultFile)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return errors.New(ErrorSuitableVaultDirNotFound)
	}
	f, err := os.OpenFile(vaultFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New(ErrorVaultFileWriteFailure)
	}
	_, err = f.WriteString(content)
	if err != nil {
		return errors.New(ErrorVaultFileWriteFailure)
	}
	defer f.Close()
	return nil
}