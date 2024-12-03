package logger

import (
	"log"
	"os"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Info(message string) {
	log.Println("INFO: " + message)
}

func Error(message string) {
	log.Printf("ERROR: %s", message)
}
