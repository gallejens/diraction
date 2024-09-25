package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var isDevBuild = false

func main() {
	workingFile, workingDir := getWorkingDir()

	if strings.Contains(workingDir, "go-build") {
		isDevBuild = true
		log.Println("[INFO] Running in development mode")
	}

	loadConfig(workingDir)

	if !isDevBuild {
		checkStartupApp(workingFile)
	}

	go startWatcher()

	// block forever
	select {}
}

func getWorkingDir() (string, string) {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return executable, filepath.Dir(executable)
}
