package a

import (
	"log"
	"os"

	core "github.com/mdelapenya/release-drafter-monorepo"
)

const p = "🛳️ [A]"

type ALogger struct {
	logger *log.Logger
}

func DoSomething(msg string) {
	core.WithLogger(&ALogger{
		logger: log.New(os.Stdout, p, log.LstdFlags),
	})

	core.Log.Warn(msg)
}

func (l *ALogger) Debug(msg string) {
	l.log("DEBUG", msg)
}

func (l *ALogger) Info(msg string) {
	l.log("INFO", msg)
}

func (l *ALogger) Warn(msg string) {
	l.log("WARN", msg)
}

func (l *ALogger) Error(msg string) {
	l.log("ERROR", msg)
}

func (l *ALogger) Fatal(msg string) {
	l.logger.SetPrefix("FATAL: ")
	l.logger.Fatalln(msg)
}

func (l *ALogger) log(prefix string, msg string) {
	l.logger.SetPrefix(p + prefix + " ")
	l.logger.Println(msg)
}
