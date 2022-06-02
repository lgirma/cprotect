package cprotect

import (
	"encoding/hex"
	"errors"
	"os/exec"
	"strings"
)

func getHDDSerialNumber() (string, error) {
	wmicCmd := exec.Command("wmic", "diskdrive", "get", "serialnumber")
	output, err := wmicCmd.Output()
	if err != nil {
		return "", err
	}
	sn := strings.Split(strings.TrimSuffix(string(output), "\n"), "\n")
	return strings.TrimSpace(sn[1]), nil
}

func getBaseBoardSerialNumber() (string, error) {
	wmicCmd := exec.Command("wmic", "baseboard", "get", "serialnumber")
	output, err := wmicCmd.Output()
	if err != nil {
		return "", err
	}
	sn := strings.Split(strings.TrimSuffix(string(output), "\n"), "\n")
	return strings.TrimSpace(sn[1]), nil
}

func GetRequestCode(productCode string) (string, error) {
	if len(productCode) == 0 {
		return "", errors.New(ErrorProductIdEmpty)
	}
	errs := make([]error, 0)
	hddId, err := getHDDSerialNumber()
	if err != nil {
		errs = append(errs, err)
	}
	mbId, err := getBaseBoardSerialNumber()
	if err != nil {
		errs = append(errs, err)
	}

	hardwareId := hddId + mbId
	if len(hardwareId) == 0 {
		if len(errs) > 0 {
			return "", errors.New(ErrorWMICExecutionFailure)
		}
		return "", errors.New(ErrorHardwareIdEmpty)
	}

	id := strings.ToUpper(hardwareId + "/" + productCode)
	return hex.EncodeToString([]byte(id)), nil
}