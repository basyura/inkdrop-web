package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Inkdrop Inkdrop `toml:"inkdrop"`
	ExecDir string
}

type Inkdrop struct {
	Root     string `tmol:"root"`
	UserID   string `toml:"userID"`
	Password string `toml:"password"`
}

func (c Config) String() string {
	s, _ := json.MarshalIndent(c, "", "  ")
	return string(s)
}

var config_ Config

func Instance() Config {
	return config_
}

func Initialize() (err error) {

	var exePath string
	exePath, err = os.Executable()
	if err != nil {
		return
	}
	execDir := filepath.Dir(exePath)

	path := filepath.Join(execDir, "config.toml")
	if _, err = toml.DecodeFile(path, &config_); err != nil {
		err = fmt.Errorf("failed to load "+path+".\ncause: %w", err)
	}
	config_.ExecDir = execDir

	return err
}
