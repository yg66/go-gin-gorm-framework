package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/yg66/go-gin-gorm-framework/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Init init logger
func Init(config *model.LogConfig) (err error) {
	// Defines a log level type pointer
	var level = new(zapcore.Level)
	err = level.UnmarshalText([]byte(config.Level))
	if err != nil {
		return
	}

	// create core
	// Development mode, log output to terminal
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	//// log output destination
	//// 1.output to file only
	//writeSyncer := getLogWriter(config.Filename, config.MaxSize, config.MaxBackups, config.MaxAge)
	//core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level)
	//// 2.Simultaneous output to file and console
	//core = zapcore.NewTee(
	//	zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	//	zapcore.NewCore(consoleEncoder, writeSyncer, level),
	//)

	// Create a logger object
	log := zap.New(core, zap.AddCaller(), zap.Development())
	defer log.Sync()
	// Replace the global logger, and then use zap.L() calls in other packages
	zap.ReplaceGlobals(log)
	return
}

func getEncoder() zapcore.Encoder {
	// Use NewProductionEncoderConfig provided by zap
	encoderConfig := zap.NewProductionEncoderConfig()
	// set time format
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// time key
	encoderConfig.TimeKey = "time"
	// level
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// Display caller information
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// Returns the log editor in console format
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	// Archiving slice logs with lumberjack
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
