package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	workingFile, workingDir := getWorkingDir()

	loadConfig(workingDir)

	checkStartupApp(workingFile)

	go startWatcher()

	select {}
}

func getWorkingDir() (string, string) {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return executable, filepath.Dir(executable)
}
