package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewConfig() *Config {
	return &Config{}
}

func NewZap(cfg *Config) *zap.Logger {
	return zap.New(
		zapcore.NewCore(getEncoder(cfg), getWriteSyncer(cfg), getLoggerLevel(cfg)),
		getOptions(cfg)...,
	)
}

func getEncoder(cfg *Config) zapcore.Encoder {
	var encoderCfg zapcore.EncoderConfig

	if cfg.Env == "development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if cfg.Env == "production" {
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if cfg.Env == "local" {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.LevelKey = "level"
		encoderCfg.NameKey = "logger"
		encoderCfg.CallerKey = "caller"
		encoderCfg.MessageKey = "message"
		encoderCfg.StacktraceKey = "stacktrace"
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeDuration = zapcore.SecondsDurationEncoder
		encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	return encoder
}

func getWriteSyncer(cfg *Config) zapcore.WriteSyncer {
	if cfg.Path == "" {
		return zapcore.Lock(os.Stdout)
	}

	file, err := os.OpenFile(cfg.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return zapcore.AddSync(file)
}

func getLoggerLevel(cfg *Config) zap.AtomicLevel {
	var level zapcore.Level
	switch cfg.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	if err := level.Set(cfg.Level); err != nil {
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	return zap.NewAtomicLevelAt(level)
}

func getOptions(cfg *Config) []zap.Option {
	if cfg.Env == "production" {
		return []zap.Option{
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		}
	}

	return []zap.Option{
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	}
}
