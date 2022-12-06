package initialize

import (
	"GinRESTful/restapi/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// StructGetLogWriter 参数个数太多，定义一个结构体传参
type StructGetLogWriter struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// InitLogger 初始化Logger
// 参数：
//		无
// 返回值：
//		无
func InitLogger() {
	lg := global.Lg
	logInfo := global.Settings.LogsInfo
	stGeWr := StructGetLogWriter{
		FileName:   logInfo.FileName,
		MaxSize:    logInfo.MaxSize,
		MaxBackups: logInfo.MaxBackups,
		MaxAge:     logInfo.MaxAge,
		Compress:   logInfo.Compress,
	}
	writeSyncer := getLogWriter(stGeWr)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(logInfo.Level))
	if err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg = zap.New(core, zap.AddCaller())
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
}

// getEncoder 获取编码器
// 参数：
//		无
// 返回值：
//		zapcore.Encoder：zap的编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter 日志写入器
// 参数：
//		stGeWr：Log编写器
// 返回值：
//		zapcore.WriteSyncer：写同步器
func getLogWriter(stGeWr StructGetLogWriter) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   stGeWr.FileName,
		MaxSize:    stGeWr.MaxSize,
		MaxBackups: stGeWr.MaxBackups,
		MaxAge:     stGeWr.MaxAge,
		Compress:   stGeWr.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
