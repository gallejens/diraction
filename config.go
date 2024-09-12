package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

//go:embed default.toml
var defaultConfig []byte

type Config struct {
	Folders []struct {
		Path string
	}
}

func loadConfig(workingDir string) {
	configPath := filepath.Join(workingDir, "config.toml")

	if !doesFileExist(configPath) {
		os.WriteFile(configPath, defaultConfig, fs.ModeAppend)
	}

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = toml.Unmarshal(configFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range cfg.Folders {
		fmt.Println(folder.Path)
	}
}
