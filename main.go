package main

import (
	// "flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	// Create intermediate directories if they don't exist
	err1 := os.MkdirAll(filepath.Dir(destiantionFile), os.ModePerm)
	if err1 != nil {
		log.Fatal("error1: ", err1);
		return
	}
	// open the destination file and write
	if checkFileExists(destiantionFile) {// if exist open
		
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
	} else {// create if do not exist
		
		destination, err := os.Create(destiantionFile) //create the destination file
		if err != nil {
			fmt.Print("here1\n");
			log.Fatal("Error:", err)
		}
		// write
		_, err = io.Copy(destination, source) //copy the contents of source to destination file
		if err != nil {
			log.Fatal("Error:", err)
		}
	}
}
func ittrateOverDir(path string, d fs.DirEntry, backUpDir string, sourceDir string){
	if d.IsDir() {
		fmt.Printf("DIRECTORY DETECTED %s\n", d.Name())
		return // skip it
	}else {
		trimmed := strings.TrimPrefix(path, sourceDir)
		copyFile(path, backUpDir + trimmed)
		fmt.Print(".")
		// fmt.Printf("%s\n", path);
		// fmt.Printf("%s\n", sourceDir);
		// fmt.Printf("%s\n", trimmed);
		// fmt.Printf("%s\n", backUpDir);
		// fmt.Printf("%s\n", backUpDir + trimmed);
	}	
}
func main() {
	// Initial Paths initaisation
	var backUpDir, sourceDir string
	sourceDir = "/Users/this_is_mjk/projects/Summer2024_projects/Pclub_VCS/Recrument_task/sourceUpDir"
	backUpDir = "/Users/this_is_mjk/projects/Summer2024_projects/Pclub_VCS/Recrument_task/backUpDir"
	os.Setenv("SOURCE_DIR", sourceDir)
	os.Setenv("BACKUP_DIR", backUpDir)

	// ittrate over the dir
	fmt.Print("working...\n");
	err := filepath.WalkDir(os.Getenv("SOURCE_DIR"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Return the error if encountered during traversal
		}
		ittrateOverDir(path, d, os.Getenv("BACKUP_DIR"), os.Getenv("SOURCE_DIR"))
		return nil
	})
	if err != nil {
		log.Fatalf("!!!!ERROR!!!!\n\nimpossible to walk directories: %s", err)
	}
	print("\nFINISHED!\n");

}
