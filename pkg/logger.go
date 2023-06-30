package pkg

import (
	"log"
	"os"
)

type Logger struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewLogger(filepath string) Logger {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("could not init log file:%v", err)
	}
	InfoLog := log.New(file, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog := log.New(file, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Logger{
		InfoLog:  InfoLog,
		ErrorLog: ErrorLog,
	}
}
