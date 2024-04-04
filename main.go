package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/gofor-little/env"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func setEnv(sourceFile string, destiantionFile string, encryptionKey string) {
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
	if encryptionKey != "" {
		// If not, write the key-value pair
		if err := env.Write("ENCRYPTION_KEY", encryptionKey, ".env", true); err != nil {
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
func copyFile(sourceFile string, destiantionFile string, encryption bool) {
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
	if encryption {
		// Reading key
		EK, err1 := env.MustGet("ENCRYPTION_KEY")
		if err1 != nil {
			log.Fatal("Please set the ENCRYPTION_KEY by -cEK")
		}
		// Reading plaintext file
		plainText, err := os.ReadFile(sourceFile)
		if err != nil {
			log.Fatalf("read file err: %v", err.Error())
		}
		// Creating block-cypher of algorithm
		block, err := aes.NewCipher([]byte(EK))
		if err != nil {
			log.Fatalf("cipher err: %v", err.Error())
		}
		// Creating GCM mode
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			log.Fatalf("cipher GCM err: %v", err.Error())
		}
		// Generating random nonce
		nonce := make([]byte, gcm.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			log.Fatalf("nonce  err: %v", err.Error())
		}
		// ecrypt file
		cipherText := gcm.Seal(nonce, nonce, plainText, nil)

		// Writing ciphertext file
		err = os.WriteFile(destiantionFile, cipherText, 0777)
		if err != nil {
			log.Fatalf("write file err: %v", err.Error())
		}
	} else {
		_, err = io.Copy(destination, source)
		if err != nil {
			log.Fatal("Error104:", err)
		}
	}
}
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
	_, err = file.WriteString(formattedTime +  " ecryption? " + fmt.Sprintf("%t", encryption) + " ")
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
		copyFile(path, backUpDir+trimmed, encryption)
		fmt.Print(".")

		// Append to log file
		_, err = file.WriteString("Copying file " + d.Name() + "\n")
		if err != nil {
			log.Fatal(err)
		}

	}
}
func decryptFile(key string, finalName string, encryptedFilePath string) {
		// Reading ciphertext file
		cipherText, err := os.ReadFile(encryptedFilePath)
		if err != nil {
			log.Fatal(err)
		}
		// Creating block of algorithm
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			log.Fatalf("cipher err: %v", err.Error())
		}
		// Creating GCM mode
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			log.Fatalf("cipher GCM err: %v", err.Error())
		}
		// Deattached nonce and decrypt
		nonce := cipherText[:gcm.NonceSize()]
		cipherText = cipherText[gcm.NonceSize():]
		plainText, err := gcm.Open(nil, nonce, cipherText, nil)
		if err != nil {
			log.Fatalf("decrypt file err: %v", err.Error())
		}
		// Writing decryption content
		err = os.WriteFile(finalName, plainText, 0777)
		if err != nil {
			log.Fatalf("write file err: %v", err.Error())
		}
}
	
func main() {
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
	setEnv(*SourcePtr, *BackUPPtr, *encrytpionKeyPtr)
	// Get the variables back
	sourceDir, backupDir := getEnv()
	// flag to print the current path and encryption key
	if *printPathPtr {
		fmt.Printf("Source Directory: %s\nBackUp Directory: %s\n", sourceDir, backupDir)
		return
	}
	if *printEKPtr {
		// Get an environment variable's value, receiving an error if it is not set or is empty.
		EK, err1 := env.MustGet("ENCRYPTION_KEY")
		if err1 != nil {
			log.Fatal("Please set the ENCRYPTION_KEY by -cEK")
		}
		fmt.Printf("Encryption Key: %s\n", EK)
		return
	}
	// decrypt
	if *decrypt {
		// key
		fmt.Print("Enter the key: ")
		var key string
		_, err := fmt.Scanf("%s\n", &key)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// final file name
		fmt.Print("Enter the final file name: ")
		var finalFileName string
		_, err = fmt.Scanf("%s\n", &finalFileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// encrypted file path
		fmt.Print("Enter the encrypted file path: ")
		var encryptedFilePath string
		_, err = fmt.Scanf("%s\n", &encryptedFilePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		decryptFile(key, finalFileName, encryptedFilePath)
		return
	}
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
