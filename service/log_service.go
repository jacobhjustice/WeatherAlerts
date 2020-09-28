package service

import (
	"log"
	"os"
	"time"

	"github.com/jacobhjustice/WeatherAlerts/model/enum"
)

type ILogService interface {
	Log(message string, severity enum.LogSeverity)
}

type LogService struct {
	ILogService
}

// Public
func (l *LogService) Log(message string, severity enum.LogSeverity) {
	file := l.openAndGetFile()
	fmsg := l.formatMessage(message, severity)
	file.WriteString(fmsg)
	l.closeFile(file)
}

// Private
func (l *LogService) closeFile(f *os.File) {
	f.Close()
}

func (l *LogService) formatMessage(message string, severity enum.LogSeverity) string {
	fmsg := ""
	fmsg += l.formatMessageTimestamp()
	fmsg += "\t"
	fmsg += string(severity)
	fmsg += "\t"
	fmsg += message

	return fmsg
}

func (l *LogService) formatMessageTimestamp() string {
	dt := time.Now()
	layout := "15:04:05"
	time := dt.Format(layout)
	return time
}

func (l *LogService) getFileName() string {
	dt := time.Now()
	layout := "2006-01-02"
	day := dt.Format(layout)
	fname := "logs/" + day
	return fname
}

func (l *LogService) openAndGetFile() *os.File {
	fileName := l.getFileName()
	fileName += ".log"
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return f
}
