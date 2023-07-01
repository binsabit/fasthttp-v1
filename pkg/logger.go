package pkg

import (
	"log"
	"os"
)

func NewLogger(filepath, prefix string) *log.Logger {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("could not init log file:%v", err)
	}
	logger := log.New(file, prefix+":\t", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}
