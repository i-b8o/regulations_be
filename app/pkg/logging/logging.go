package logging

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

// writerHook is a hook that writes logs of specified LogLevels to specified Writer
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	// format log entry to string
	line, err := entry.String()
	if err != nil {
		log.Fatal(err)
	}
	// write it to all writers
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
	}
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (l *Logger) LWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

func (l *Logger) LWithFields(fields map[string]interface{}) *Logger {
	return &Logger{l.WithFields(fields)}
}

func Init(level string) {
	// take a string level and returns the Logrus log level constant.
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatalln(err)
	}

	// Creates a new logger
	l := logrus.New()
	// Set a ReportCaller field
	l.SetReportCaller(true)

	// format log into text
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := path.Base(f.File)
			return fmt.Sprintf("%s:%d", fileName, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	// Send all logs to nowhere by default
	l.SetOutput(ioutil.Discard)
	// Send all logs to Stdout
	l.AddHook(&writerHook{Writer: []io.Writer{os.Stdout}, LogLevels: logrus.AllLevels})

	// set the logger level
	l.SetLevel(logrusLevel)

	e = logrus.NewEntry(l)
}
