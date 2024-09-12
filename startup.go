package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var STARTUP_FOLDER = `\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\`
var shortcutPath string

func checkStartupApp(workingFile string) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	shortcutPath = filepath.Join(userHomeDir, STARTUP_FOLDER, "diraction.lnk")

	if doesFileExist(shortcutPath) {
		fmt.Println("This app is already in the startup folder")
		return
	}

	err = createShortcut(workingFile, shortcutPath)
	if err != nil {
		log.Fatal(err)
	}
}

func createShortcut(executablePath, shortcutPath string) error {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return err
	}
	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer wshell.Release()
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", shortcutPath)
	if err != nil {
		return err
	}
	idispatch := cs.ToIDispatch()
	oleutil.PutProperty(idispatch, "TargetPath", executablePath)
	oleutil.CallMethod(idispatch, "Save")
	return nil
}

func removeShortcut() {
	err := os.Remove(shortcutPath)
	if err != nil {
		log.Fatal(err)
	}
}
