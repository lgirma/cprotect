package cprotect

import (
	"encoding/hex"
	"errors"
	"strings"
)

func GetRequestCode(productCode string, fingerprintService FingerprintService) (string, error) {
	elevatedPrivilege := isCurrentUserRoot()
	if len(productCode) == 0 {
		return "", errors.New(ErrorProductIdEmpty)
	}
	hardwareId, err := fingerprintService.GetMachineId(elevatedPrivilege)
	if err != nil {
		return "", err
	}
	id := strings.ToUpper(hardwareId + "/" + productCode)
	return hex.EncodeToString([]byte(id)), nil
}