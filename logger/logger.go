package logger

import (
"path"
"time"

rotatelogs "github.com/lestrrat-go/file-rotatelogs"
log "github.com/sirupsen/logrus"
)

type LogConfigs struct {
	LogPath    string
	FileName   string
	MaxAge     int
	LogLevel   int
	PerLogSize int64
	MaxCount   int
}

// 0：PanicLevel 1：FatalLevel，2：ErrorLevel，3：WarnLevel，4：InfoLevel，5：DebugLevel，6：TraceLevel
// InitLog 初始化一个全局的log, 初始化之后可以直接使用 log "github.com/sirupsen/logrus"
func InitLog(conf *LogConfigs) error {

	baseLogPath := path.Join(conf.LogPath, conf.FileName)
	//maxAge := time.Duration(conf.MaxAge*24) * time.Hour
	rotationTime := 24 * time.Hour

	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d.log",
		//rotatelogs.WithLinkName(baseLogPath),

		// 设置日志分割的时间
		rotatelogs.WithRotationTime(rotationTime),

		rotatelogs.WithRotationSize(conf.PerLogSize),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		// rotatelogs.WithMaxAge(time.Hour*24),
		//rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationCount(uint(conf.MaxCount)),
	)

	if err != nil {
		log.Errorf("config local file system for logger error:%v", err)
		return err
	}

	// 设置输出
	log.SetOutput(writer)

	// 设置日志级别
	log.SetLevel(log.Level(conf.LogLevel))

	// 设置日志格式
	log.SetFormatter(&log.TextFormatter{})

	//log.SetReportCaller(true)
	return nil
}

