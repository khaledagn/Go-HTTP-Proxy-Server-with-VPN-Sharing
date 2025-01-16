package logger

import (
	"log"
	"os"
)

var loggerInstance *log.Logger

// initializes the custom logger
func InitLogger() {
	loggerInstance = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	loggerInstance.Println("Logger initialized")
}

// info logs informational messages
func Info(message string) {
	if loggerInstance == nil {
		InitLogger() 
	}
	loggerInstance.SetPrefix("INFO: ")
	loggerInstance.Println(message)
}

// error logs error messages
func Error(err error) {
	if loggerInstance == nil {
		InitLogger() 
	}
	loggerInstance.SetPrefix("ERROR: ")
	loggerInstance.Println(err.Error())
}

// warn logs warning messages
func Warn(message string) {
	if loggerInstance == nil {
		InitLogger()  
	}
	loggerInstance.SetPrefix("WARN: ")
	loggerInstance.Println(message)
}

// debug logs debug messages
func Debug(message string) {
	if loggerInstance == nil {
		InitLogger() 
	}
	loggerInstance.SetPrefix("DEBUG: ")
	loggerInstance.Println(message)
}
