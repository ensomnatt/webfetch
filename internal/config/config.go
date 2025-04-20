package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AppDir string
	AppFile string
	Port string
}

type userConfig struct {
	Port string
}

func NewConfig() *Config {
	config := &Config{}

	config.getAppDir()
	config.getAppFile()

	config.getUserConfig()

	return config
}

func (c *Config) getAppDir() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		logrus.Errorf("error with getting user config directory: %v", err)
		return
	}

	c.AppDir = filepath.Join(configDir, "webfetch")

	_, err = os.Stat(c.AppDir)
	if os.IsNotExist(err) {
		logrus.Info("didnt find app directory. creating...")
		err := os.Mkdir(c.AppDir, 0755)
		if err != nil {
			logrus.Errorf("error with creating app dir: %v", err)
			return
		}
	}
}

func (c *Config) getAppFile() {
	exampleCfg := `
	port = ":6969"
	`

	c.AppFile = filepath.Join(c.AppDir, "config.toml")

	_, err := os.Stat(c.AppFile)
	if os.IsNotExist(err) {
		logrus.Info("didnt find config file. creating...")
		file, err := os.Create(c.AppFile)
		if err != nil {
			logrus.Errorf("error with creating config.toml: %v", err)
			return
		}

		_, err = file.Write([]byte(exampleCfg))
		if err != nil {
			logrus.Errorf("error with writing example config file: %v", err)
			return
		}
	}
}

func (c *Config) getUserConfig() {
	data, err := os.ReadFile(c.AppFile)
	if err != nil {
		logrus.Errorf("error with reading config.toml: %v", err)
		return
	}

	userCfg := &userConfig{}
	err = toml.Unmarshal(data, userCfg)
	if err != nil {
		logrus.Errorf("error with unmarshaling config.toml: %v", err)
		return
	}

	logrus.Info("read config file")

	c.Port = userCfg.Port
} 
