package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func startWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	for _, cfgFolder := range cfg.Folders {
		registerFolder(watcher, cfgFolder.Path, cfgFolder.IncludeSubfolders)
	}

	select {}
}

func registerFolder(watcher *fsnotify.Watcher, path string, includeSubfolders bool) {
	fileInfo, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) || !fileInfo.IsDir() {
		log.Printf("[WARNING] %s is not a directory\n", path)
		return
	}

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}

	if !includeSubfolders {
		return
	}

	dirContent, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, dirEntry := range dirContent {
		if !dirEntry.IsDir() {
			continue
		}

		registerFolder(watcher, filepath.Join(path, dirEntry.Name()), includeSubfolders)
	}
}
