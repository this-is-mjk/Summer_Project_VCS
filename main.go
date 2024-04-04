package main

import (
	// "flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
func copyFile(sourceFile string, destiantionFile string) {
	// open the source file
	source, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer source.Close()
	// open the destination file and write
	if checkFileExists(destiantionFile) {
		// if exist open
		destination, err := os.Open(destiantionFile)
		if err != nil {
			log.Fatal("Error:", err)
		}
		defer source.Close()
		// write
		_, err = io.Copy(destination, source) //copy the contents of source to destination file
		if err != nil {
			log.Fatal("Error:", err)
		}
	} else {
		// create if not exist
		destination, err := os.Create(destiantionFile) //create the destination file
		if err != nil {
			log.Fatal("Error:", err)
		}
		// write
		_, err = io.Copy(destination, source) //copy the contents of source to destination file
		if err != nil {
			log.Fatal("Error:", err)
		}
	}
}
func ittrateOverDir(path string, d fs.DirEntry, err error, backUpDir string, sorceDir string) error {
	if d.IsDir() {
		fmt.Println("DIRECTORY DETECTED %s", d.Name())
		return nil
	}
	copyFile(path, )
}
func main() {
	// Initial Paths initaisation
	var backUpDir, sourceDir string
	sourceDir = "/Users/this_is_mjk/projects/Summer2024_projects/Pclub_VCS/Recrument_task/sourceUpDir/text1Test.txt"
	backUpDir = "/Users/this_is_mjk/projects/Summer2024_projects/Pclub_VCS/Recrument_task/backUpDir/text1Test.txt"
	os.Setenv("SOURCE_DIR", sourceDir)
	os.Setenv("BACKUP_DIR", backUpDir)

	fmt.Println("Let's Start")
	// copyFile(os.Getenv("SOURCE_DIR"), os.Getenv("BACKUP_DIR"))
	// ittrate over the dir
	err := filepath.WalkDir(os.Getenv("BACKUP_DIR"), func(path string, d fs.DirEntry, err error) error {
		return ittrateOverDir(path, d, err, os.Getenv("BACKUP_DIR"), os.Getenv("SOURCE_DIR"))
	})
	if err != nil {
		log.Fatalf("!!!!ERROR!!!!\n\nimpossible to walk directories: %s", err)
	}

}
