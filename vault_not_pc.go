//+build wasm

package cprotect

import "errors"

func IsInstalled(product string, password string) (bool, error) {
	return false, errors.New("NOT_IMPLEMENTED")
}

func IsActivationCodeValid(password string, requestCode string, activationCode string) bool {
	return false
}

func Install(product string, activationCode string) error {
	return errors.New("NOT_IMPLEMENTED")
}