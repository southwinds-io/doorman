package core

import (
	"net/url"
	"southwinds.dev/artisan/core"
)

type RetryLogger struct {
}

func (l *RetryLogger) Error(msg string, kv ...interface{}) {
	u := kv[5].(*url.URL)
	core.ErrorLogger.Printf("%s, %s %s%s", msg, kv[1], u.Host, u.Path)
}

func (l *RetryLogger) Info(msg string, kv ...interface{}) {
	u := kv[3].(*url.URL)
	core.InfoLogger.Printf("%s, %s %s%s", msg, kv[1], u.Host, u.Path)
}

func (l *RetryLogger) Debug(msg string, kv ...interface{}) {
	if core.InDebugMode() {
		u := kv[3].(*url.URL)
		core.DebugLogger.Printf("%s, %s %s%s", msg, kv[1], u.Host, u.Path)
	}
}

func (l *RetryLogger) Warn(msg string, kv ...interface{}) {
	u := kv[3].(*url.URL)
	core.ErrorLogger.Printf("%s, %s %s%s", msg, kv[1], u.Host, u.Path)
}
