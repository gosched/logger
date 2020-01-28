package logger

import (
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

var (
	file *os.File
	err  error
)

// Init .
func Init(logPath string) error {
	file, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	Info = log.New(io.MultiWriter(os.Stdout, file), `Info  `, log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, file), `Error `, log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// Close .
func Close() error {
	return file.Close()
}
