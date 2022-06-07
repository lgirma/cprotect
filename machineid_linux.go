//+build !windows

package cprotect

import "errors"

func GetMachineId() (string, error) {
	return "", errors.New("NOT_IMPLEMENTED")
}