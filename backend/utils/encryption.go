package utilities

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"

	"golang.org/x/crypto/blake2b"
)

func CreateHash(password string) string {
	hash, err := blake2b.New(32, nil)
	if err != nil {
		log.Println(err)
	}
	_, err = hash.Write([]byte(password))
	if err != nil {
		log.Println(err)
	}
	return string(hash.Sum(nil))
}

func GenerateAESKey() []byte {
	key := make([]byte, 32)
	_, _ = rand.Read(key)
	return key
}

func Encrypt(data, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Println(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return plaintext
}
