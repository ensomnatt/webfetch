package config

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type Config struct {
	AppDir string
}

func NewConfig() *Config {
	configDir, err := os.UserConfigDir()
	if err != nil {
		logrus.Errorf("error with getting user config directory: %v", err)
		return nil
	}
	appDir := filepath.Join(configDir, "webfetch")

	_, err = os.Stat(appDir)
	if os.IsNotExist(err) {
		logrus.Info("didnt find app directory. creating...")
		err := os.Mkdir(appDir, 0755)
		if err != nil {
			logrus.Errorf("error with creating app dir: %v", err)
			return nil
		}
	}

	return &Config{
		AppDir: appDir,
	}
}
