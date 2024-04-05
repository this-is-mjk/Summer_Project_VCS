package encrypterhelper

import (
	"github.com/gofor-little/env"
	"log"
	"os"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)
func Encrypter(sourceFile string, destiantionFile string) {
	// Reading key
	EK, err1 := env.MustGet("ENCRYPTION_KEY")
	if err1 != nil {
		log.Fatal("Please set the ENCRYPTION_KEY by -cEK")
	}
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
}
func DecryptFile(key string, finalName string, encryptedFilePath string) {
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