package model

import (
	"gorm.io/gorm"
	"time"
)

// ====== config ======

type Config struct {
	BasicsConfig *BasicsConfig `mapstructure:"basics" json:"basics"`
	DbConfig     *DbConfig     `mapstructure:"db" json:"db"`
	LogConfig    *LogConfig    `mapstructure:"log" json:"log"`
}

type BasicsConfig struct {
	IsDev bool `mapstructure:"is_dev" json:"is_dev"`
	Port  uint `mapstructure:"port" json:"port"`
}

type DbConfig struct {
	MysqlDns string `mapstructure:"mysql_dns" json:"mysql_dns"`
	LogLevel string `mapstructure:"log_level" json:"log_level"`
}

type LogConfig struct {
	Level      string `mapstructure:"level" json:"level"`
	OutFile    bool   `mapstructure:"out_file" json:"out_file"`
	Filename   string `mapstructure:"filename" json:"filename"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"`
}

// ====== mysql table ======

type Cache struct {
	gorm.Model
	CacheKey   string `gorm:"unique;not null"`
	CacheValue string
	Expired    time.Time
}
