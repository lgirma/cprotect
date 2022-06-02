package main

import (
	"cprotect"
	"fmt"
	"golang.org/x/term"
	"os"
)

func main() {
	product := ""
	reqCode := ""
	reqCodeThisPC := ""
	activationCode := ""
	password := ""

	fmt.Println("CProtect Admin 1.0")
	fmt.Print("Enter Product: ")
	fmt.Scanln(&product)
	reqCodeThisPC, err := cprotect.GetRequestCode(product)
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
	isInstalled, err := cprotect.IsInstalled(product, password)
	if err != nil {
		fmt.Println("Is Installed? !Failure: " + err.Error())
	} else {
		fmt.Println(fmt.Sprintf("Is Installed? %v", isInstalled))
	}
	fmt.Print("Enter Request Code: ")
	fmt.Scanln(&reqCode)

	enc, err := cprotect.Encrypt(reqCode, password)
	if err != nil {
		panic(err)
	}
	dec, err := cprotect.Decrypt(enc, password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Request Code: %s\n", reqCode)
	fmt.Printf("Encryped: %s\n", enc)
	fmt.Printf("Decrypted: %s\n", dec)

	fmt.Print("Activation Code to Install (empty to skip): ")
	fmt.Scanln(&activationCode)
	if len(activationCode) > 0 {
		err = cprotect.Install(product, activationCode)
		if err != nil {
			fmt.Println("Failure: " + err.Error())
		} else {
			fmt.Println("Successfully installed in vault.")
		}
	} else {
		fmt.Println("Skipped installation")
	}

}