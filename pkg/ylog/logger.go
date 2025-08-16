package ylog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
	"youras/infra/config"
)

var (
	globalLogger *zap.SugaredLogger
	accessLogger *zap.SugaredLogger
	once         sync.Once
)

func InitProductionLogger(serviceName string, cfg *config.LogConfig) {
	once.Do(func() {
		// 1. 定义日志级别
		atomicLevel := zap.NewAtomicLevel()
		atomicLevel.SetLevel(zapcore.Level(cfg.Level)) // 从环境变量读取

		// 2. 编码器配置（生产环境用JSON）
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.MessageKey = "message"
		encoderConfig.TimeKey = "timestamp"

		cfg.Path = filepath.Join(cfg.Path, serviceName)
		accessLogger = initAccessLogger(cfg, serviceName, encoderConfig)
		globalLogger = initGlobalLogger(cfg, serviceName, encoderConfig, atomicLevel)

	})
}

// 初始化专用访问日志Logger
func initAccessLogger(cfg *config.LogConfig, serviceName string, encoderConfig zapcore.EncoderConfig) *zap.SugaredLogger {
	// 访问日志使用JSON格式
	accessEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// 访问日志文件配置
	accessWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(cfg.Path, "access.log"),
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})

	// 访问日志核心：只记录Info级别
	accessCore := zapcore.NewCore(
		accessEncoder,
		accessWriter,
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel
		}),
	)

	// 构建访问日志Logger（不记录调用者/堆栈）
	logger := zap.New(accessCore,
		zap.AddCallerSkip(2),                    // 跳过中间件层
		zap.AddStacktrace(zapcore.InvalidLevel), // 禁用堆栈
	)

	return logger.Sugar().With(
		"service", serviceName,
		"type", "access",
	)
}

// 初始化全局应用Logger
func initGlobalLogger(cfg *config.LogConfig, serviceName string, encoderConfig zapcore.EncoderConfig, atomicLevel zap.AtomicLevel) *zap.SugaredLogger {
	var cores []zapcore.Core

	// 控制台输出（开发环境）
	if cfg.ConsoleDebug {
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stdout),
			atomicLevel,
		)
		cores = append(cores, consoleCore)
	}

	// 应用日志文件（所有级别）
	appEncoder := zapcore.NewJSONEncoder(encoderConfig)
	appWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(cfg.Path, "app.log"),
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})
	appCore := zapcore.NewCore(
		appEncoder,
		appWriter,
		atomicLevel,
	)
	cores = append(cores, appCore)

	// 错误日志文件（Warn+级别）
	errorEncoder := zapcore.NewJSONEncoder(encoderConfig)
	errorWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(cfg.Path, "error.log"),
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})
	errorCore := zapcore.NewCore(
		errorEncoder,
		errorWriter,
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.WarnLevel
		}),
	)
	cores = append(cores, errorCore)

	// 组合核心
	combinedCore := zapcore.NewTee(cores...)

	// 构建全局Logger
	logger := zap.New(combinedCore,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	return logger.Sugar().With(
		"service", serviceName,
		"type", "app",
	)
}

// 获取全局应用Logger
func Log() *zap.SugaredLogger {
	return globalLogger
}

// 获取访问专用Logger
func Access() *zap.SugaredLogger {
	return accessLogger
}
