package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
  AppDir string
}

func NewConfig() *Config {
  configDir, err := os.UserConfigDir()
  if err != nil {
    _ = fmt.Errorf("error with getting user config directory: %v", err)
    return nil
  }
  appDir := filepath.Join(configDir, "webfetch")

  _, err = os.Stat(appDir)
  if os.IsNotExist(err) {
    err := os.Mkdir(appDir, 0755)
    if err != nil {
      _ = fmt.Errorf("error with creating app dir: %v", err)
      return nil
    }
  }

  return &Config{
    AppDir: appDir,
  }
}
