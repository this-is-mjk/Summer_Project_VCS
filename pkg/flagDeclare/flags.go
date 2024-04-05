package flags

import (
	encrypt "Summer_Project_VCS/pkg/encrypterHelper"
	envHelper "Summer_Project_VCS/pkg/envHelper"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofor-little/env"
)

func DeclareFlags() (string, string, *bool) {
	// flags
	SourcePtr := flag.String("cSP", "", "To set source directory path, include -cSP=")
	BackUPPtr := flag.String("cBP", "", "To set backUp directory path, include -cBP=")
	encrytpionKeyPtr := flag.String("cEK", "", "To set encryption key, include -cEK=")
	printPathPtr := flag.Bool("Printpaths", false, "To print the current set paths, include -Printpaths")
	printEKPtr := flag.Bool("PrintEK", false, "To print the current set Key, include -PrintEK")
	encryption := flag.Bool("E", false, "To encrypt and back-it-up, include -E")
	decrypt := flag.Bool("D", false, "To decrypt, include -D")
	flag.Parse()

	// Initial Paths initaisation
	envHelper.SetEnv(*SourcePtr, *BackUPPtr, *encrytpionKeyPtr)
	// Get the variables back
	sourceDir, backupDir := envHelper.GetEnv()
	// flag to print the current path and encryption key
	if *printPathPtr {
		fmt.Printf("Source Directory: %s\nBackUp Directory: %s\n", sourceDir, backupDir)
		os.Exit(-1)
	}
	if *printEKPtr {
		// Get an environment variable's value, receiving an error if it is not set or is empty.
		EK, err1 := env.MustGet("ENCRYPTION_KEY")
		if err1 != nil {
			log.Fatal("Please set the ENCRYPTION_KEY by -cEK")
		}
		fmt.Printf("Encryption Key: %s\n", EK)
		os.Exit(-1)
	}
	// decrypt
	if *decrypt {
		// key
		fmt.Print("Enter the key: ")
		var key string
		_, err := fmt.Scanf("%s\n", &key)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		// final file name
		fmt.Print("Enter the final file name: ")
		var finalFileName string
		_, err = fmt.Scanf("%s\n", &finalFileName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		// encrypted file path
		fmt.Print("Enter the encrypted file path: ")
		var encryptedFilePath string
		_, err = fmt.Scanf("%s\n", &encryptedFilePath)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		encrypt.DecryptFile(key, finalFileName, encryptedFilePath)
		os.Exit(-1)
	}
	return backupDir, sourceDir, encryption
}
