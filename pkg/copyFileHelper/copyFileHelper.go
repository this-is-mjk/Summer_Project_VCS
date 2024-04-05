package copyfilehelper

import (
	"io"
	"log"
	"os"
	"path/filepath"
	encrypt "Summer_Project_VCS/pkg/encrypterHelper"
)

func CopyFile(sourceFile string, destiantionFile string, encryption bool) {
	// Create intermediate directories if they don't exist
	err1 := os.MkdirAll(filepath.Dir(destiantionFile), os.ModePerm)
	if err1 != nil {
		log.Fatal("Error102: ", err1)
		return
	}
	// now copy them
	if encryption { // ecrypt it if with flag -E, it will encrypt and write
		encrypt.Encrypter(sourceFile, destiantionFile)
	} else {
		// Open the source file
		source, err := os.Open(sourceFile)
		if err != nil {
			log.Fatal("Error101:", err)
		}
		defer source.Close()
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
}
