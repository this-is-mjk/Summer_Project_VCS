package main

import (
	"flag"
	"fmt"
	"github.com/gofor-little/env"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func setEnv(sourceFile string, destiantionFile string) {
	// Get the current working directory
	cwd, _ := os.Getwd()
	// Load an .env file and set the key-value pairs as environment variables.
	if err := env.Load(filepath.Join(cwd, ".env")); err != nil {
		log.Fatal("Error201:", err)
	}
	// Check if the key-value pairs already exist in the .env file
	if sourceFile != "" {
		// If not, write the key-value pair
		if err := env.Write("SORCE_PATH", sourceFile, ".env", true); err != nil {
			log.Fatal("Error202:", err)
		}
	}
	if destiantionFile != "" {
		// If not, write the key-value pair
		if err := env.Write("DESTINATION_PATH", destiantionFile, ".env", true); err != nil {
			log.Fatal("Error203:", err)
		}
	}
}
func getEnv() (string, string) {
	// Get the current working directory
	cwd, _ := os.Getwd()
	// Load an .env file and set the key-value pairs as environment variables.
	if err := env.Load(filepath.Join(cwd, ".env")); err != nil {
		log.Fatal("Error301:", err)
	}
	// Get an environment variable's value, receiving an error if it is not set or is empty.
	sourceFile, err1 := env.MustGet("SORCE_PATH")
	if err1 != nil {
		// log.Fatal("Error302:", err1)
		log.Fatal("Please set the SOURCE_PATH by -cSP")
	}
	destiantionFile, err2 := env.MustGet("DESTINATION_PATH")
	if err2 != nil {
		// log.Fatal("Error303:", err2)
		log.Fatal("Please set the DESTINATION_PATH by -cBP")
	}
	return sourceFile, destiantionFile
}
func copyFile(sourceFile string, destiantionFile string) {
	// Open the source file
	source, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal("Error101:", err)
	}
	defer source.Close()

	// Create intermediate directories if they don't exist
	err1 := os.MkdirAll(filepath.Dir(destiantionFile), os.ModePerm)
	if err1 != nil {
		log.Fatal("Error102: ", err1)
		return
	}
	// Open the destination file for writing, create it if it doesn't exist
	destination, err := os.Create(destiantionFile)
	if err != nil {
		log.Fatal("Error103:", err)
	}
	defer destination.Close()
	// Write the contents of source to the destination file
	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatal("Error104:", err)
	}
}
func ittrateOverDir(path string, d fs.DirEntry, backUpDir string, sourceDir string) {
	if d.IsDir() {
		fmt.Printf("DIRECTORY DETECTED %s\n", d.Name())
		return // skip it
	} else {
		trimmed := strings.TrimPrefix(path, sourceDir)
		copyFile(path, backUpDir+trimmed)
		fmt.Print(".")
	}
}
func main() {
	SourcePtr := flag.String("cSP", "", "To set source directory path")
	BackUPPtr := flag.String("cBP", "", "To set backUp directory path")
	printPathPtr := flag.Bool("PrintPath", false, "To print the current set paths")
	
	flag.Parse()
	// Initial Paths initaisation
	setEnv(*SourcePtr, *BackUPPtr)
	// Get the env variables
	sourceDir, backupDir := getEnv()
	// flag to print the current path
	if *printPathPtr {
		fmt.Printf("Source Directory: %s\nBackUp Directory: %s\n", sourceDir, backupDir);
		return
	}
	// ittrate over the dir
	fmt.Print("working...\n")
	err := filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Return the error if encountered during traversal
		}
		ittrateOverDir(path, d, backupDir, sourceDir)
		return nil
	})
	if err != nil {
		log.Fatalf("!!!!ERROR!!!!\n\nimpossible to walk directories: %s", err)
	}

	print("\nFINISHED!\n")

}
