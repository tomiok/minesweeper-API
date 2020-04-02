package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log   *zap.Logger
	sugar *zap.SugaredLogger
)

/*
InitDefault Function initializes a logger using uber-go/zap package in the application.
*/
func InitDefault(env string) {
	disableCaller := true
	if env != "prod" {
		disableCaller = false
	}
	conf := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "file",
			MessageKey:   "msg",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		DisableCaller:    disableCaller,
	}
	log, _ = conf.Build()
	sugar = log.Sugar()
}

func Log() *zap.Logger {
	_ = log.Sync()
	return log
}

func Sugar() *zap.SugaredLogger {
	_ = sugar.Sync()
	return sugar
}
