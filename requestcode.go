package cprotect

import (
	"encoding/hex"
	"errors"
	"strings"
)

func GetRequestCode(productCode string) (string, error) {
	if len(productCode) == 0 {
		return "", errors.New(ErrorProductIdEmpty)
	}
	hardwareId, err := GetMachineId()
	if err != nil {
		return "", err
	}
	id := strings.ToUpper(hardwareId + "/" + productCode)
	return hex.EncodeToString([]byte(id)), nil
}