package cprotect

import (
	"log"
	"os/exec"
	"os/user"
	"strings"
)

func runCmd(program string, args ...string) (string, error) {
	cmdInstance := exec.Command(program, args...)
	cmdOutputBytes, err := cmdInstance.Output()
	if err != nil {
		return "", err
	}
	return string(cmdOutputBytes), nil
}

func isCurrentUserRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Printf("[isRoot] unable to get current user: %s", err)
		return false
	}
	return strings.TrimSpace(currentUser.Username) == "root"
}
