package logger

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

func init() {
	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.DebugLevel)

	baseLogPath := path.Join(viper.GetString("log.path"))
	rotatelogs.WithMaxAge(time.Hour * 24)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H",
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName(baseLogPath),

		// WithRotationTime设置日志分割的时间,这里设置为24小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),

		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// WithRotationCount设置文件清理前最多保存的个数.
		rotatelogs.WithMaxAge(time.Hour*24*7),
		//rotatelogs.WithRotationCount(maxRemainCnt),
	)
	if err != nil {
		panic(err)
	}
	//
	pathMap := lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}
	log.AddHook(lfshook.NewHook(pathMap, &log.TextFormatter{}))
}
