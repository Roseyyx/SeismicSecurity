package utilities

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
)

func createHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func encrypt(data []byte, passwordHash string) []byte {
	block, _ := aes.NewCipher([]byte(passwordHash))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passwordHash string) []byte {
	block, _ := aes.NewCipher([]byte(passwordHash))
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return plaintext
}

func EncryptFile(password string, file *os.File) {
	// Read the file
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	file.Read(buffer)

	// Encrypt the file
	encryptedData := encrypt(buffer, createHash(password))

	// Write the encrypted data to the file
	file.Seek(0, 0)
	file.Write(encryptedData)
}

func DecryptFile(password string, file *os.File) {
	// Read the file
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	file.Read(buffer)

	// Decrypt the file
	decryptedData := decrypt(buffer, createHash(password))

	// Write the decrypted data to the file
	file.Seek(0, 0)
	file.Write(decryptedData)
}
