package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type Logger interface {
	Info(format string, values ...interface{})
	Warn(format string, values ...interface{})
	Error(format string, values ...interface{})
	Fatal(format string, values ...interface{})
}

type logger struct {
	WriteToConsole     bool
	WriteToFile        bool
	infoLoggerFile     *log.Logger
	infoLoggerConsole  *log.Logger
	warnLoggerFile     *log.Logger
	warnLoggerConsole  *log.Logger
	errorLoggerFile    *log.Logger
	errorLoggerConsole *log.Logger
	fatalLoggerFile    *log.Logger
	fatalLoggerConsole *log.Logger
}

func NewLogger(writeToConsole bool, writeToFile bool) (Logger, error) {
	var file *os.File = nil
	if writeToFile {
		if _, err := os.Stat("./logs"); os.IsNotExist(err) {
			if err = os.Mkdir("./logs", os.ModePerm); err != nil {
				return nil, err
			}
		}

		currentDateStamp := time.Now().Format("02-01-2006")

		var err error
		file, err = os.OpenFile(fmt.Sprintf("./logs/%s.log", currentDateStamp), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
	}

	infoColor := color.New(color.BgBlack).Add(color.FgHiWhite).Add(color.Bold)
	infoLoggerFile := log.New(file, "INFO  "+" ", log.Ldate|log.Ltime)
	infoLoggerConsole := log.New(os.Stdout, infoColor.Sprint(" INFO  ")+" ", log.Ldate|log.Ltime)

	warnColor := color.New(color.BgBlack).Add(color.FgHiYellow).Add(color.Bold)
	warnLoggerFile := log.New(file, "WARN  "+" ", log.Ldate|log.Ltime)
	warnLoggerConsole := log.New(os.Stdout, warnColor.Sprint(" WARN  ")+" ", log.Ldate|log.Ltime)

	errorColor := color.New(color.BgBlack).Add(color.FgHiRed).Add(color.Bold)
	errorLoggerFile := log.New(file, "ERROR "+" ", log.Ldate|log.Ltime)
	errorLoggerConsole := log.New(os.Stdout, errorColor.Sprint(" ERROR ")+" ", log.Ldate|log.Ltime)

	fatalColor := color.New(color.BgHiRed).Add(color.FgBlack).Add(color.Bold)
	fatalLoggerFile := log.New(file, "FATAL "+" ", log.Ldate|log.Ltime)
	fatalLoggerConsole := log.New(os.Stdout, fatalColor.Sprint(" FATAL ")+" ", log.Ldate|log.Ltime)

	return &logger{
		WriteToConsole:     writeToConsole,
		WriteToFile:        writeToFile,
		infoLoggerConsole:  infoLoggerConsole,
		infoLoggerFile:     infoLoggerFile,
		warnLoggerConsole:  warnLoggerConsole,
		warnLoggerFile:     warnLoggerFile,
		errorLoggerConsole: errorLoggerConsole,
		errorLoggerFile:    errorLoggerFile,
		fatalLoggerConsole: fatalLoggerConsole,
		fatalLoggerFile:    fatalLoggerFile,
	}, nil
}

func (l *logger) Info(format string, values ...interface{}) {
	if l.WriteToConsole {
		l.infoLoggerConsole.Printf(l.getCallPath()+": "+format, values...)
	}

	if l.WriteToFile {
		l.infoLoggerFile.Printf(l.getCallPath()+": "+format, values...)
	}
}

func (l *logger) Warn(format string, values ...interface{}) {
	if l.WriteToConsole {
		l.warnLoggerConsole.Printf(l.getCallPath()+": "+format, values...)
	}

	if l.WriteToFile {
		l.warnLoggerFile.Printf(l.getCallPath()+": "+format, values...)
	}
}

func (l *logger) Error(format string, values ...interface{}) {
	if l.WriteToConsole {
		l.errorLoggerConsole.Printf(l.getCallPath()+": "+format, values...)
	}

	if l.WriteToFile {
		l.errorLoggerFile.Printf(l.getCallPath()+": "+format, values...)
	}
}

func (l *logger) Fatal(format string, values ...interface{}) {
	if l.WriteToConsole {
		l.fatalLoggerConsole.Printf(l.getCallPath()+": "+format, values...)
	}

	if l.WriteToFile {
		l.fatalLoggerFile.Printf(l.getCallPath()+": "+format, values...)
	}

	os.Exit(1)
}

func (l *logger) getCallPath() string {
	_, file, line, _ := runtime.Caller(2)

	split := strings.Split(file, "/")
	splitLen := len(split)

	if splitLen >= 3 {
		return fmt.Sprintf("%s:%d", strings.Join(split[splitLen-3:], "/"), line)
	} else if splitLen >= 2 {
		return fmt.Sprintf("%s:%d", strings.Join(split[splitLen-2:], "/"), line)
	}

	return fmt.Sprintf("%s:%d", strings.Join(split[splitLen-1:], "/"), line)
}
