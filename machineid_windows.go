//+build windows

package cprotect

import (
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

func getDiskDriveId(elevatedPrivilege bool) (string, error) {
	output, err := queryWMIC("diskdrive", "get", "serialnumber") //exec.Command("wmic", "diskdrive", "get", "serialnumber")
	if err != nil {
		return "", err
	}
	sn := strings.Split(strings.TrimSuffix(output, "\n"), "\n")
	return strings.TrimSpace(sn[1]), nil
}

func getMotherboardId(elevatedPrivilege bool) (string, error) {
	output, err := queryWMIC("baseboard", "get", "serialnumber") //exec.Command("wmic", "diskdrive", "get", "serialnumber")
	if err != nil {
		return "", err
	}
	sn := strings.Split(strings.TrimSuffix(output, "\n"), "\n")
	return strings.TrimSpace(sn[1]), nil
}