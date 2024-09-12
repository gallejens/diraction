package main

import (
	"log"

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
		err = watcher.Add(cfgFolder.Path)
		if err != nil {
			log.Fatal(err)
		}
	}

	select {}
}
