package main

import (
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// logger := zap.NewExample()
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/foo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	// atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// atom.SetLevel(zap.ErrorLevel)

	core := zapcore.NewCore(
		// zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.NewJSONEncoder(encoderCfg),
		w,
		zap.InfoLevel,
	)

	logger := zap.New(core)
	defer logger.Sync()

	// logger := zap.New(zapcore.NewCore(
	// 	zapcore.NewJSONEncoder(encoderCfg),
	// 	zapcore.Lock(os.Stdout),
	// 	atom,
	// ))
	// defer logger.Sync()

	url := "http://example.org/api"

	logger.Info(
		"failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugar := logger.Sugar()
	sugar.Infow(
		"failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
