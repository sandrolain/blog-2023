package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/mlkem"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	// Generate a new decapsulation key (equivalent to a private key)
	decKey, err := mlkem.GenerateKey1024()
	if err != nil {
		panic("failed to generate decapsulation key")
	}

	// Extract the encapsulation key (equivalent to a public key)
	encKey := decKey.EncapsulationKey()

	fmt.Printf("decKey: %v\n", decKey)
	fmt.Printf("encKey: %v\n", encKey)

	// Define a plaintext message
	msg := []byte("This is a test message")

	// Generate a new shared key and relative ciphertext
	sharedKey, sharedKeyCiphertext := encKey.Encapsulate()

	// sharedKey is a 256 bit symmetric key that can be used for AES encryption

	var msgCyphertext []byte

	// Encrypt with AES GCM 256
	{
		// Create a block cipher
		block, err := aes.NewCipher(sharedKey)
		if err != nil {
			panic("failed to create cipher block")
		}

		// Create a GCM cipher
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			panic("failed to create GCM cipher")
		}

		// Generate a random nonce for each encryption
		nonce := make([]byte, gcm.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			panic("failed to generate nonce")
		}

		// Seal will append the ciphertext to the nonce
		// The nonce is prepended to the ciphertext and will be used for decryption
		msgCyphertext = gcm.Seal(nonce, nonce, msg, nil)
	}

	fmt.Printf("msgCyphertext: %v\n", msgCyphertext)

	sharedKey, e := decKey.Decapsulate(sharedKeyCiphertext)
	if e != nil {
		panic("failed to decapsulate key")
		return
	}

	// Create a block cipher
	block, err := aes.NewCipher(sharedKey)
	if err != nil {
		panic("failed to create block cipher")
	}

	// Create a GCM cipher
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("failed to create GCM cipher")
	}

	// Extract the nonce from the ciphertext
	nonceSize := gcm.NonceSize()
	if len(msgCyphertext) < nonceSize {
		panic("ciphertext is too short")
	}

	// Extract the nonce from the ciphertext
	nonce := msgCyphertext[:nonceSize]
	if len(nonce) != nonceSize {
		panic("invalid nonce size")
	}

	// Extract the actual ciphertext after the nonce
	encryptedData := msgCyphertext[nonceSize:]
	if len(encryptedData) < 1 {
		panic("ciphertext is too short")
	}

	msg, err = gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		panic("failed to decrypt message")
	}

	fmt.Printf("msg: %v\n", string(msg))

}
