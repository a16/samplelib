package samplelib

import (
	"log"
	"os"
	"sync"
)

var Logger = log.New(os.Stderr, "", log.LstdFlags)
var LogMu sync.Mutex

func SetLogger(l *log.Logger) {
	if l == nil {
		l = log.New(os.Stderr, "", log.LstdFlags)
	}
	logMu.Lock()
	logger = l
	logMu.Unlock()
}

func logf(format string, v ...interface{}) {
	logMu.Lock()
	log.Printf(format, v...)
	logMu.Unlock()
}
