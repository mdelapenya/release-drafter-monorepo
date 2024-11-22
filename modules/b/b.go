package b

import (
	"log"
	"os"

	core "github.com/mdelapenya/release-drafter-monorepo"
)

const p = "[B] "

type BLogger struct {
	logger *log.Logger
}

func DoSomething(msg string) {
	core.WithLogger(&BLogger{
		logger: log.New(os.Stdout, p, log.LstdFlags),
	})

	core.Log.Info(msg)
}

func (l *BLogger) Debug(msg string) {
	l.log("DEBUG", msg)
}

func (l *BLogger) Info(msg string) {
	l.log("INFO", msg)
}

func (l *BLogger) Warn(msg string) {
	l.log("WARN", msg)
}

func (l *BLogger) Error(msg string) {
	l.log("ERROR", msg)
}

func (l *BLogger) Fatal(msg string) {
	l.logger.SetPrefix("FATAL: ")
	l.logger.Fatalln(msg)
}

func (l *BLogger) log(prefix string, msg string) {
	l.logger.SetPrefix(p + prefix + ": ")
	l.logger.Println(msg)
}
