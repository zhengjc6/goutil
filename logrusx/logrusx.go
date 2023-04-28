package logrusx

import (
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type LogrusConfig struct {
	Level string `yaml:"Level"`
	Path  string `yaml:"Path"`
	Save  uint   `yaml:"Save"` //保存文件数，每天一个文件
}

var (
	Log  = logrus.New()
	once sync.Once
)

func InitLog(cfg LogrusConfig) {
	once.Do(func() {
		Log.SetOutput(os.Stdout)
		var loglevel logrus.Level
		err := loglevel.UnmarshalText([]byte(cfg.Level))
		if err != nil {
			fmt.Printf("set log level fail%v", err)
			panic(err)
		}
		Log.SetLevel(loglevel)
		Log.SetFormatter(&logrus.TextFormatter{})
		LocalFilesystemLogger(cfg.Path, cfg.Save)
	})
}

func logWriter(logPath string, level string, save uint) *rotatelogs.RotateLogs {
	logFullPath := path.Join(logPath, level)
	logwriter, err := rotatelogs.New(
		logFullPath+".%Y%m%d.log",
		rotatelogs.WithLinkName(logFullPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(save),        // 文件最大保存份数
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		panic(err)
	}
	return logwriter
}

func LocalFilesystemLogger(logPath string, save uint) {
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logWriter(logPath, "debug", save), // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  logWriter(logPath, "info", save),
		logrus.WarnLevel:  logWriter(logPath, "warn", save),
		logrus.ErrorLevel: logWriter(logPath, "error", save),
		logrus.FatalLevel: logWriter(logPath, "fatal", save),
		logrus.PanicLevel: logWriter(logPath, "panic", save),
	}, &logrus.JSONFormatter{})
	Log.AddHook(lfHook)
}
