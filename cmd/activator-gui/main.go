package main

import (
	"errors"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lgirma/cprotect"
)

var Product string
var Password string

var vault_service cprotect.VaultService
var fingerprint_service cprotect.FingerprintService

func main() {

	if Product == "" {
		Product = "MyProduct"
	}
	if Password == "" {
		Password = "MyPassword"
	}

	a := app.New()
	win := a.NewWindow(Product + " Activator 2.0")
	win.Resize(fyne.Size{Width: 600, Height: 400})
	icon_file, err := os.ReadFile(filepath.Join("res", "logo.png"))
	if err != nil {
		log.Printf("failed to load icon: %v", err)
	} else {
		win.SetIcon(fyne.NewStaticResource("icon", icon_file))
	}

	// activationCode := ""
	fingerprint_service = cprotect.GetFingerprintService()
	vault_service = cprotect.GetVaultService()
	forAllUsers := false
	passwordBytes := []byte(Password)
	password := string(passwordBytes[:])

	req_code, err := cprotect.GetRequestCode(Product, fingerprint_service)
	if err != nil {
		dialog.ShowError(errors.New("failed to get request code"), win)
		return
	}
	is_installed, err := cprotect.IsInstalled(Product, password, forAllUsers,
		fingerprint_service, vault_service)
	if err != nil {
		dialog.ShowError(errors.New("failed to check installation"), win)
	}

	title_lbl := canvas.NewText(Product+" Activator", color.White)
	title_lbl.TextStyle = fyne.TextStyle{Bold: true}
	title_lbl.Alignment = fyne.TextAlignCenter
	title_lbl.TextSize = 20

	lbl_product := canvas.NewText("Product:", color.RGBA{255, 255, 255, 128})
	input_product := widget.NewLabel(Product)
	input_product.TextStyle = fyne.TextStyle{}
	box_product := container.New(layout.NewHBoxLayout(), lbl_product, input_product)

	lbl_installed := canvas.NewText("Installed?", color.RGBA{255, 255, 255, 128})
	input_installed := widget.NewLabel("No")
	if is_installed {
		input_installed.SetText("Yes")
	}
	box_installed := container.New(layout.NewHBoxLayout(), lbl_installed, input_installed)

	check_forall := widget.NewCheck("For all users", func(for_all bool) {
		input_installed.SetText("Checking...")
		is_installed, err := cprotect.IsInstalled(Product, password, for_all,
			fingerprint_service, vault_service)
		if err != nil {
			dialog.ShowError(errors.New("failed to check installation"), win)
			input_installed.SetText("failed to check")
		} else if is_installed {
			input_installed.SetText("Yes")
		} else {
			input_installed.SetText("No")
		}
	})

	lbl_reqcode := canvas.NewText("Request Code:", color.RGBA{255, 255, 255, 128})
	input_reqcode := widget.NewEntry()
	input_reqcode.SetText(req_code)
	input_reqcode.MultiLine = true
	input_reqcode.Wrapping = fyne.TextWrapBreak
	box_reqcode := container.New(layout.NewVBoxLayout(),
		lbl_reqcode, input_reqcode)

	lbl_actcode := canvas.NewText("Activation Code:", color.RGBA{255, 255, 255, 128})
	input_actcode := widget.NewEntry()
	input_actcode.SetText("")
	input_actcode.MultiLine = true
	input_actcode.Wrapping = fyne.TextWrapBreak
	box_actcode := container.New(layout.NewVBoxLayout(),
		lbl_actcode, input_actcode)

	btn_install := widget.NewButton("Install", func() {
		install(input_actcode.Text, req_code, check_forall.Checked, win)
	})
	if is_installed {
		btn_install.Disable()
		btn_install.SetText("Installed")
	}
	btn_close := widget.NewButton("Close", func() {
		win.Close()
	})
	hbox := container.New(layout.NewHBoxLayout(),
		btn_install,
		btn_close)

	win.SetContent(container.New(layout.NewVBoxLayout(),
		title_lbl,
		box_product,
		box_installed,
		check_forall,
		box_reqcode,
		box_actcode,
		layout.NewSpacer(),
		hbox,
	))

	win.ShowAndRun()
}

func install(activation_code string, req_code string, for_all_users bool, win fyne.Window) {
	if len(activation_code) > 0 {
		if !cprotect.IsActivationCodeValid(Password, req_code, activation_code) {
			dialog.ShowError(errors.New("invalid activation code"), win)
			return
		}
		err := cprotect.Install(Product, activation_code, for_all_users, vault_service)
		if err != nil {
			dialog.ShowError(errors.New("installation failure: "+err.Error()), win)
		} else {
			dialog.ShowInformation("Success", "Successfully installed", win)
			win.Close()
			return
		}
	} else {
		dialog.ShowInformation("Missing Activation Code", "Enter activation code to install", win)
	}
}
