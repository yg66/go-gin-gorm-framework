package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BasicsConfig *BasicsConfig `yaml:"basics"`
	DbConfig     *DbConfig     `yaml:"db"`
	LogConfig    *LogConfig    `yaml:"log"`
}

type BasicsConfig struct {
	IsDev      bool `yaml:"isDev"`
	Port       uint `yaml:"port"`
	StackTrace bool `yaml:"stackTrace"`
}

type DbConfig struct {
	MysqlDns string `yaml:"mysqlDns"`
	LogLevel string `yaml:"logLevel"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	OutFile    bool   `yaml:"outFile"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"maxSize"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
}

func InitConfigByYml() (iConfig *Config, err error) {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return
	}

	iConfig = &Config{}
	err = yaml.Unmarshal(data, iConfig)
	return
}
