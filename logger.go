/*
   @Author:huolun
   @Date:2021/7/19
   @Description
*/
package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
)

const LAYOUT = "2006-01-02 15:04:05.123"

type ZLogger struct {
	*zap.Logger
}

func RFC3339TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, LAYOUT, enc)
}
func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}

func NewZLogger(develop, formatJSON bool, traceID, spanID string) *ZLogger {
	var ZapLogger *zap.Logger
	logPath := "stdout.log"
	hook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    128,     // 每个日志文件保存的大小 单位:M
		MaxAge:     7,       // 文件最多保存多少天
		MaxBackups: 30,      // 日志文件最多保存多少个备份
		Compress:   false,   // 是否压缩
	}
	var levelEncoder zapcore.LevelEncoder
	atomicLevel := zap.NewAtomicLevel()
	var writes []zapcore.WriteSyncer
	if develop {
		levelEncoder = zapcore.LowercaseColorLevelEncoder
		atomicLevel.SetLevel(zap.DebugLevel)
		writes = append(writes, zapcore.AddSync(os.Stdout))
	} else {
		writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
		levelEncoder = zapcore.LowercaseLevelEncoder
		atomicLevel.SetLevel(zap.InfoLevel)
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "datetime",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    levelEncoder,
		EncodeTime:     RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	var core zapcore.Core
	if formatJSON {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(writes...),
			atomicLevel,
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(writes...),
			atomicLevel,
		)
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	field := zap.Fields(zap.String("trace_id", traceID), zap.String("span_id", spanID))
	// 构造日志
	ZapLogger = zap.New(core, caller, development, field)

	return &ZLogger{
		Logger: ZapLogger,
	}
}
