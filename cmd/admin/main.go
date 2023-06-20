package main

import (
	"fmt"
	"os"

	"github.com/lgirma/cprotect"
	"golang.org/x/term"
)

func pressAnyKeyToExit() {
	fmt.Print("Press [ENTER] to exit...")
	fmt.Scanln()
}

const ColorReset = "\033[0m"
const ColorGreen = "\033[32m"
const ColorRed = "\033[31m"

func showError(msg string) {
	fmt.Print(ColorRed + msg + ColorReset)
}

func showSuccess(msg string) {
	fmt.Print(ColorGreen + msg + ColorReset)
}

func main() {
	product := ""
	reqCode := ""
	reqCodeThisPC := ""
	activationCode := ""
	password := ""
	fingerprintService := cprotect.GetFingerprintService()
	vaultService := cprotect.GetVaultService()
	forAllUsers := false

	fmt.Println("CProtect Admin 1.0")
	fmt.Print("Enter Product: ")
	fmt.Scanln(&product)
	reqCodeThisPC, err := cprotect.GetRequestCode(product, fingerprintService)
	if err != nil {
		fmt.Println("Req Code (this PC): !Failure: " + err.Error())
	} else {
		fmt.Println("Req Code (this PC): " + reqCodeThisPC)
	}
	fmt.Print("Enter Password: ")
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Print("[!Failed to hide password: " + err.Error() + "] ")
		fmt.Scan(&password)
	} else {
		password = string(passwordBytes[:])
	}
	fmt.Print("\n")
	isInstalled, err := cprotect.IsInstalled(product, password, forAllUsers,
		fingerprintService, vaultService)
	if err != nil {
		showError("\nFailed to check for existing installation: " + err.Error())
	} else {
		if isInstalled {
			showSuccess("Software already installed\n")
		} else {
			fmt.Printf("Software not installed")
		}

	}
	fmt.Print("Enter Request Code: ")
	fmt.Scanln(&reqCode)

	enc, err := cprotect.Encrypt(reqCode, password)
	if err != nil {
		showError(err.Error())
		pressAnyKeyToExit()
		return
	}
	dec, err := cprotect.Decrypt(enc, password)
	if err != nil {
		showError(err.Error())
		pressAnyKeyToExit()
		return
	}
	fmt.Printf("Request Code: %s\n", reqCode)
	fmt.Printf("Encryped: %s\n", enc)
	fmt.Printf("Decrypted: %s\n", dec)

	fmt.Print("Activation Code to Install (empty to skip): ")
	fmt.Scanln(&activationCode)
	if len(activationCode) > 0 {
		err = cprotect.Install(product, activationCode, forAllUsers, vaultService)
		if err != nil {
			showError("Failure: " + err.Error())
		} else {
			showSuccess("Successfully installed in vault.")
		}
	} else {
		fmt.Println("Skipped installation")
	}
	pressAnyKeyToExit()
}
