package c

import (
	"log"
	"os"

	core "github.com/mdelapenya/release-drafter-monorepo"
)

const p = "🚗 [C]"

type CLogger struct {
	logger *log.Logger
}

func DoSomething(msg string) {
	core.WithLogger(&CLogger{
		logger: log.New(os.Stdout, p, log.LstdFlags),
	})

	core.Log.Error(msg)
}

func (l *CLogger) Debug(msg string) {
	l.log("DEBUG", msg)
}

func (l *CLogger) Info(msg string) {
	l.log("INFO", msg)
}

func (l *CLogger) Warn(msg string) {
	l.log("WARN", msg)
}

func (l *CLogger) Error(msg string) {
	l.log("ERROR", msg)
}

func (l *CLogger) Fatal(msg string) {
	l.logger.SetPrefix("FATAL: ")
	l.logger.Fatalln(msg)
}

func (l *CLogger) log(prefix string, msg string) {
	l.logger.SetPrefix(p + prefix + " ")
	l.logger.Println(msg)
}
