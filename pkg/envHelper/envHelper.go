package envHelper 
import (
	"os"
	"github.com/gofor-little/env"
	"path/filepath"
	"log"
)
func SetEnv(sourceFile string, destiantionFile string, encryptionKey string) {
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
func GetEnv() (string, string) {
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