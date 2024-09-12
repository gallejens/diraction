package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	workingFile, workingDir := getWorkingDir()
	fmt.Println(workingDir)

	loadConfig(workingDir)

	checkStartupApp(workingFile)

	go func() {
		for {
			fmt.Println("Hello from goroutine")
			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}

func getWorkingDir() (string, string) {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return executable, filepath.Dir(executable)
}
