package logx

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func Test_zap(t *testing.T) {
	zapSugar()
	zapCommon()
}

// 不指明类型用的是反射
func zapSugar() {
	url := "ww.baidu.com"
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

// 确定类型，性能更优
func zapCommon() {
	url := "ww.baidu.com"
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
