package main

import (
	"cprotect"
	"encoding/hex"
	"fmt"
)

var Product string
var Password string

const ColorReset = "\033[0m"
const ColorGreen = "\033[32m"
const ColorRed = "\033[31m"

func showError(msg string) {
	fmt.Print(ColorRed + "[x ERROR] " + msg + ColorReset)
}

func showSuccess(msg string) {
	fmt.Print(ColorGreen + "[v SUCCESS] " + msg + ColorReset)
}

func pressAnyKeyToExit() {
	fmt.Print("Press [ENTER] to exit...")
	fmt.Scanln()
}

func main() {
	reqCodeThisPC := ""
	activationCode := ""
	passwordBytes, _ := hex.DecodeString(Password)
	password := string(passwordBytes[:])

	fmt.Println("CProtect Activator 1.0")
	if len(Product) > 0 {
		fmt.Println("Product: " + Product)
	} else {
		fmt.Print("Enter Product: ")
		fmt.Scan(&Product)
	}
	reqCodeThisPC, err := cprotect.GetRequestCode(Product)
	if err != nil {
		showError("Failed to get request code\n")
		pressAnyKeyToExit()
		return
	} else {
		fmt.Println("Req Code (this PC): " + reqCodeThisPC)
	}
	isInstalled, err := cprotect.IsInstalled(Product, password)
	if err != nil {
		showError("Failed to check installation: " + err.Error() + "\n")
	} else if isInstalled {
		showSuccess("Software already activated\n")
	}

	fmt.Print("Activation Code (ENTER to skip): ")
	fmt.Scanln(&activationCode)
	if len(activationCode) > 0 {
		if !cprotect.IsActivationCodeValid(password, reqCodeThisPC, activationCode) {
			showError("Invalid activation code\n")
			pressAnyKeyToExit()
			return
		}
		err = cprotect.Install(Product, activationCode)
		if err != nil {
			showError("Installation failure: " + err.Error() + "\n")
		} else {
			showSuccess("Successfully installed in vault.")
		}
	} else {
		fmt.Println("Skipped installation")
	}
	pressAnyKeyToExit()
}
