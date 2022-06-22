package model

import (
	"gorm.io/gorm"
	"time"
)

// ====== config ======

type Config struct {
	BasicsConfig    *BasicsConfig    `mapstructure:"basics" json:"basics"`
	DbConfig        *DbConfig        `mapstructure:"db" json:"db"`
	ReCAPTCHAConfig *ReCAPTCHAConfig `mapstructure:"reCAPTCHA" json:"reCAPTCHA"`
	JobConfig       *JobConfig       `mapstructure:"job" json:"job"`
	SmtpConfig      *SmtpConfig      `mapstructure:"smtp" json:"smtp"`
	LogConfig       *LogConfig       `mapstructure:"log" json:"log"`
}

type BasicsConfig struct {
	IsDev         bool   `mapstructure:"is_dev" json:"is_dev"`
	IsCoreApiTest bool   `mapstructure:"is_core_api_test" json:"is_core_api_test"`
	Port          uint   `mapstructure:"port" json:"port"`
	JwtSecret     string `mapstructure:"jwt_secret" json:"jwt_secret"`
	KeyFilePath   string `mapstructure:"key_file_path" json:"key_file_path"`
	ViewGateway   string `mapstructure:"view_gateway" json:"view_gateway"`
	TempSecret    string `mapstructure:"temp_secret" json:"temp_secret"`
}

type DbConfig struct {
	MysqlDns string `mapstructure:"mysql_dns" json:"mysql_dns"`
}

type ReCAPTCHAConfig struct {
	Secret string `mapstructure:"secret" json:"secret"`
}

type JobConfig struct {
	JobCycle string `mapstructure:"job_cycle" json:"job_cycle"`
}

type SmtpConfig struct {
	FromEmail       string `mapstructure:"from_email" json:"from_email"`
	FromAccount     string `mapstructure:"from_account" json:"from_account"`
	FromEmailPasswd string `mapstructure:"from_email_passwd" json:"from_email_passwd"`
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
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
