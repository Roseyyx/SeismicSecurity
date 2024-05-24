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

func CreateHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func Encrypt(data []byte, passwordHash string) []byte {
	block, _ := aes.NewCipher([]byte(passwordHash))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte, passwordHash string) []byte {
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
	encryptedData := Encrypt(buffer, CreateHash(password))

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
	decryptedData := Decrypt(buffer, CreateHash(password))

	// Write the decrypted data to the file
	file.Seek(0, 0)
	file.Write(decryptedData)
}
