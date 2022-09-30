package main

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"log"
	"os/exec"
)

func main() {
	systray.Run(onReady, onExit)
}

func changeGateway(gw string) {
	cmd := exec.Command("route", "delete", "0.0.0.0")
	_ = cmd.Run()
	cmd = exec.Command("route", "add", "0.0.0.0", "mask", "0.0.0.0", gw)
	_ = cmd.Run()
}

func onReady() {
	systray.SetIcon(icon.Data)
	mDefault := systray.AddMenuItem("Default", "192.168.1.1")
	mProxy := systray.AddMenuItem("Proxy", "192.168.1.23")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	for {
		select {
		case <-mDefault.ClickedCh:
			log.Println("changing to 192.168.1.1")
			go changeGateway("192.168.1.1")
		case <-mProxy.ClickedCh:
			log.Println("changing to 192.168.1.23")
			go changeGateway("192.168.1.23")
		case <-mQuit.ClickedCh:
			log.Println("exiting")
			systray.Quit()
		}
	}
}

func onExit() {

}
