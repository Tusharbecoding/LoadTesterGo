package logger

import (
	"log"
)

func InitLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func LogError(workerID int, err error) {
	log.Printf("[Worker %d] Error: %s\n", workerID, err.Error())
}

func LogInfo(workerID int, message string) {
	log.Printf("[Worker %d] Info: %s\n", workerID, message)
}
