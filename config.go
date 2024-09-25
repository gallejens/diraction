package main

import (
	_ "embed"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

//go:embed default.toml
var defaultConfig []byte

type Config struct {
	Settings struct {
		PlaySound bool
	}
	Folders []struct {
		Path              string
		IncludeSubfolders bool
	}
}

var cfg Config

func loadConfig(workingDir string) {
	configPath := filepath.Join(workingDir, "config.toml")

	if !doesFileExist(configPath) {
		os.WriteFile(configPath, defaultConfig, fs.ModeAppend)
	}

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal(configFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[INFO] Config loaded")
}
