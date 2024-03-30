// Package logger provides a custom-configured zap.Logger.
package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevels is a struct that holds the log level and whether to add stacktrace or not.
type LogLevels struct {
	LogLevel      zapcore.Level
	AddStacktrace bool
}

// NewLogger creates and returns a custom-configured zap.Logger.
func NewLogger(lv *LogLevels) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	options := []zap.Option{
		zap.AddCaller(),
	}

	logLevel, stacktrace := getLogLevelFromEnvOrDefault(lv)
	if stacktrace {
		options = append(options, zap.AddStacktrace(logLevel))
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		zap.NewAtomicLevelAt(logLevel),
	)

	logger := zap.New(core, options...)

	return logger
}

func getLogLevelFromEnvOrDefault(lv *LogLevels) (zapcore.Level, bool) {
	if lv != nil {
		return lv.LogLevel, lv.AddStacktrace
	}

	envLevel := zapcore.InfoLevel

	envLevelStr := os.Getenv("LOG_LEVEL")
	if envLevelStr != "" {
		err := envLevel.UnmarshalText([]byte(envLevelStr))
		if err != nil {
			log.Printf("Invalid LOG_LEVEL '%s', using default InfoLevel. Error: %v", envLevelStr, err)
		}
	}

	envStacktrace := false

	envStacktraceBool := os.Getenv("STACKTRACE")
	if envStacktraceBool == "true" {
		envStacktrace = true
	}

	return envLevel, envStacktrace
}
