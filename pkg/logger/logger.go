package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func LogWrier(logToWrite string, backUpDir string) {
	// Create intermediate directories if they don't exist
	err1 := os.MkdirAll(filepath.Dir(backUpDir+"/log.txt"), os.ModePerm)
	if err1 != nil {
		log.Fatal("Error102: ", err1)
		return
	}
	// Open the log file in append mode
	file, err := os.OpenFile(backUpDir+"/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// write
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	_, err = file.WriteString(formattedTime + " " + logToWrite)
	if err != nil {
		log.Fatal(err)
	}
}
