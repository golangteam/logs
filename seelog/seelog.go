package seelog

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cihub/seelog"
	"github.com/golangteam/function/file"
)

const (
	logConfig = `<seelog type="asynctimer" asyncinterval="5000000" minlevel="debug">
	<outputs formatid="main">
		<console/>
		<rollingfile type="size" filename="__log_url__" maxsize="1024000" maxrolls="10" />
	</outputs>
	<formats>
		<format id="main" format="%Date(2006-01-02 15:04:05) [%Level] %RelFile line:%Line %Msg%n"/>
	</formats>
</seelog>`
)

//InitDefault init seelog
//
// @param pathLog    string  path of logs
// @param pathConfig string  path of log.xml
func InitDefault() {
	Init("", "")
}

//Info log info
func Info(v ...interface{}) {
	seelog.Info(v...)
}

//Init init seelog
//
// @param pathLog    string  path of logs
// @param pathConfig string  path of log.xml
func Init(logPath, configPath string) {

	if logPath == "" {
		logPath = "logs/roll.log"
	}
	tmp := filepath.Dir(logPath)
	if file.FileIsNotExist(tmp) {
		os.MkdirAll(tmp, 0764)
	}

	if configPath == "" {
		configPath = "log.xml"
	}
	tmp = filepath.Dir(configPath)
	if file.FileIsNotExist(tmp) {
		os.MkdirAll(tmp, 0764)
	}
	if file.FileIsNotExist(configPath) {
		ioutil.WriteFile(configPath, []byte(strings.Replace(logConfig, "__log_url__", logPath, 1)), 0764)
	}
	if logger, err := seelog.LoggerFromConfigAsFile(configPath); err == nil {
		seelog.ReplaceLogger(logger)
	}
}

//Flush flush the seelog
func Flush() {
	seelog.Flush()
}
