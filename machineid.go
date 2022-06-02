package cprotect

import (
	"encoding/hex"
	"errors"
	"os/exec"
	"strings"
	"syscall"
)

func queryWMIC(args ...string) (string, error) {
	cmdPath := "cmd.exe"
	cmdArgs := []string{"/c", "wmic"}
	for _, arg := range args {
		cmdArgs = append(cmdArgs, arg)
	}
	cmdInstance := exec.Command(cmdPath, cmdArgs...)
	cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmdOutput, err := cmdInstance.Output()
	if err != nil {
		return "", err
	}
	return string(cmdOutput), nil
}

func getHDDSerialNumber() (string, error) {
	output, err := queryWMIC("diskdrive", "get", "serialnumber") //exec.Command("wmic", "diskdrive", "get", "serialnumber")
	if err != nil {
		return "", err
	}
	sn := strings.Split(strings.TrimSuffix(output, "\n"), "\n")
	return strings.TrimSpace(sn[1]), nil
}

func getBaseBoardSerialNumber() (string, error) {
	output, err := queryWMIC("baseboard", "get", "serialnumber") //exec.Command("wmic", "diskdrive", "get", "serialnumber")
	if err != nil {
		return "", err
	}
	sn := strings.Split(strings.TrimSuffix(output, "\n"), "\n")
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