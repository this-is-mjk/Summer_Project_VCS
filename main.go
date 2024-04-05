package main

import (
	flags "Summer_Project_VCS/pkg/flagDeclare"
	copyier "Summer_Project_VCS/pkg/copyFileHelper"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ittrateOverDir(path string, d fs.DirEntry, backUpDir string, sourceDir string, encryption bool) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
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
	// Append to log file
	_, err = file.WriteString(formattedTime + " ecryption? " + fmt.Sprintf("%t", encryption) + " ")
	if err != nil {
		log.Fatal(err)
	}

	if d.IsDir() {
		log.Printf("DIRECTORY DETECTED %s\n", d.Name())
		// Append to log file
		_, err = file.WriteString("DIRECTORY DETECTED " + d.Name() + "\n")
		if err != nil {
			log.Fatal(err)
		}
		return // skip it
	} else {
		trimmed := strings.TrimPrefix(path, sourceDir)
		copyier.CopyFile(path, backUpDir+trimmed, encryption)
		fmt.Print(".")

		// Append to log file
		_, err = file.WriteString("Copying file " + d.Name() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
func main() {
	// use flags
	backupDir, sourceDir, encryption := flags.DeclareFlags()
	// ittrate over the dir
	log.Print("working...\n")
	err := filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Return the error if encountered during traversal
		}
		ittrateOverDir(path, d, backupDir, sourceDir, *encryption)
		return nil
	})
	if err != nil {
		log.Fatalf("!!!!ERROR!!!!\n\nimpossible to walk directories: %s", err)
	}

	print("\nFINISHED!\n")

}
