package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// * Creates the logs dir
func New() *LogDir {
	err := os.Mkdir(LOGS_PATH, 0666)
	if err != nil {
		return nil
	}
	return &LogDir{
		LogDirectory: LOGS_PATH,
	}
}

// * Sets the format for log files naming squeme
func SetLogFile() *os.File {
	year, month, day := time.Now().Date()
	fileName := fmt.Sprintf("%v-%v-%v.log", day, month.String(), year)
	filePath, _ := os.OpenFile(LOGS_PATH+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	return filePath
}



// * All of these use "Reciver Functions"
// Limit from what objects they can be called from. It cannot be called by itself anymore.
// * Case for Normal Logging
func (l *LogDir) Info() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// * Case for Warning Logging
func (l *LogDir) Warning() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// * Case for Error Logging
func (l *LogDir) Error() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// * Case for Fatal Error Logging
func (l *LogDir) Fatal() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *LogDir) Debug() *log.Logger {
	getFilePath := SetLogFile()
	return log.New(getFilePath, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func logging() *LogDir {
	logger := New()
	logger.Debug().Println("Log Service Started.")
	logger.Debug().Println("logging() called.")
	return logger
}
