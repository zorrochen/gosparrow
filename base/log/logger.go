package log

import (
	"github.com/cihub/seelog"
	"time"
)

var Log iLog

type iLog interface {
	EnableTimeStamp()
	ReloadConfig() error
	Close()
	Debug(format string, params ...interface{})
	Info(format string, params ...interface{})
	Warn(format string, params ...interface{}) error
	Error(format string, params ...interface{}) error
	Critical(format string, params ...interface{}) error
}

type logger struct {
	TimeStampFlag bool
	CfgFilePath   string
}

func LogInit(logCfgPath string) {
	l := &logger{TimeStampFlag: false, CfgFilePath: logCfgPath}
	err := initSeelog(l.CfgFilePath)
	if err != nil {
		panic(err)
	}

	Log = l
	return
}

func (self *logger) EnableTimeStamp() {
	self.TimeStampFlag = true
}

func (self *logger) ReloadConfig() error {
	return initSeelog(self.CfgFilePath)
}

func (self *logger) Close() {
	seelog.Flush()
}

func (self *logger) Debug(format string, args ...interface{}) {
	if self.TimeStampFlag {
		v := append([]interface{}{time.Now().Format("01-02 15:04:05.000000")}, args...)
		seelog.Debugf("%s "+format, v...)
	} else {
		seelog.Debugf(format, args...)
	}
}

func (self *logger) Info(format string, args ...interface{}) {
	if self.TimeStampFlag {
		v := append([]interface{}{time.Now().Format("01-02 15:04:05.000000")}, args...)
		seelog.Infof("%s "+format, v...)
	} else {
		seelog.Infof(format, args...)
	}
}

func (self *logger) Warn(format string, args ...interface{}) error {
	if self.TimeStampFlag {
		v := append([]interface{}{time.Now().Format("01-02 15:04:05.000000")}, args...)
		return seelog.Warnf("%s "+format, v...)
	} else {
		return seelog.Warnf(format, args...)
	}
}

func (self *logger) Error(format string, args ...interface{}) error {
	if self.TimeStampFlag {
		v := append([]interface{}{time.Now().Format("01-02 15:04:05.000000")}, args...)
		return seelog.Errorf("%s "+format, v...)
	} else {
		return seelog.Errorf(format, args...)
	}
}

func (self *logger) Critical(format string, args ...interface{}) error {
	if self.TimeStampFlag {
		v := append([]interface{}{time.Now().Format("01-02 15:04:05.000000")}, args...)
		return seelog.Criticalf("%s "+format, v...)
	} else {
		return seelog.Criticalf(format, args...)
	}
}

func initSeelog(cfgfile string) error {
	lg, err := seelog.LoggerFromConfigAsFile(cfgfile)
	if err != nil {
		return err
	}

	lg.SetAdditionalStackDepth(1)
	seelog.ReplaceLogger(lg)
	return nil
}
