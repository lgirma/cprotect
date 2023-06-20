//go:build !windows
// +build !windows

package cprotect

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func getDiskDriveId(elevatedPrivilege bool) (string, error) {
	cmdOutput, err := runCmd("lsblk", "-dnio", "KNAME,TYPE,SERIAL,RO")
	if err != nil {
		return "", err
	}
	outputStrLines := strings.Split(cmdOutput, "\n")
	for _, line := range outputStrLines {
		cols := strings.Fields(line)
		if cols[1] == "disk" && cols[3] == "0" {
			return cols[2], nil
		}
	}
	return "", errors.New("no suitable read-write disk drive found")
}

func getMotherboardId(elevatedPrivilege bool) (string, error) {
	if elevatedPrivilege {
		res, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_serial")
		if err != nil {
			return "", err
		}
		return string(res), nil
	} else {
		board_name, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_name")
		if err != nil {
			return "", err
		}
		board_version, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_version")
		if err != nil {
			return "", err
		}
		board_vendor, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_vendor")
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s-%s-%s",
			string(board_name), string(board_vendor), string(board_version)), nil
	}
}