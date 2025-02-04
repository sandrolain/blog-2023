+++
title = "Quantum-resistant asymmetric encryption with Golang, ML-KEM and AES"
date = "2025-02-04T15:42:10Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/mathew-schwartz-sb7RUrRMaC4-unsplash.jpg"
tags = ["golang", "ml-kem"]
keywords = ["golang", "ml-kem"]
description = "Here is how to use the crypto/mlkem package to encapsulate session keys with the ML-KEM encryption method"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

With the release of Go 1.24 (currently in RC), the standard library introduces [`crypto/mlkem`](https://pkg.go.dev/crypto/mlkem), implementing the ML-KEM (Kyber) key encapsulation mechanism, a quantum-resistant cryptographic primitive standardized in [NIST FIPS 203](https://csrc.nist.gov/pubs/fips/203/final).  

ML-KEM is a post-quantum key encapsulation method based on lattice-based cryptography. Unlike traditional public-key encryption schemes, which may be vulnerable to quantum attacks, ML-KEM provides a secure way to exchange symmetric encryption keys even in the presence of adversaries equipped with quantum computing capabilities.  

AES-256 remains considered secure in a post-quantum world, as its key size is large enough to resist quantum attacks such as Grover’s algorithm, which would only halve its effective security level.  

This addition marks a significant step in Go’s cryptographic ecosystem, enabling developers to future-proof applications by adopting post-quantum encryption techniques with minimal effort.  

In the following example, we will demonstrate how to use the `crypto/mlkem` package to securely encapsulate a key for AES-256 encryption.

### Example of creation of a decapsulation key and an encapsulation key

The `GenerateKey1024` method from the `crypto/mlkem` package generates a new decapsulation key (equivalent to a private key).
From the decapsulation key we can extract the encapsulation key (equivalent to a public key).

```go
// Generate a new decapsulation key (equivalent to a private key)
decKey, err := mlkem.GenerateKey1024()
if err != nil {
  panic("failed to generate decapsulation key")
}

// Extract the encapsulation key (equivalent to a public key)
encKey := decKey.EncapsulationKey()

```

### Example of encryption using encapsulation key

The `Encapsulate` method takes no input and returns a shared key and its relative ciphertext.
The generated `sharedKey` is a 256 bit symmetric key that can be used for AES encryption and the `sharedKeyCiphertext` is the relative ciphertext encapsulating the shared key.

```go

// Define a plaintext message
msg := []byte("This is a test message")

// Generate a new shared key and relative ciphertext
sharedKeyCiphertext, sharedKey := encKey.Encapsulate()

// sharedKey is a 256 bit symmetric key that can be used for AES encryption

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
  msgCyphertext := gcm.Seal(nonce, nonce, msg, nil)
}

```

### Example of decryption using decapsulation key

The encrypted message in `msgCyphertext` can be shared with the recipient, along with the encrypted key `sharedKeyCiphertext`.
The recipient, in order to decrypt the message, must have the decapsulation key `decKey` from which the encapsulation key used to generate `sharedKeyCiphertext` was derived.

```go

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
if len(cipherText) < nonceSize {
  panic("ciphertext is too short")
}

// Extract the nonce from the ciphertext
nonce := cipherText[:nonceSize]
if len(nonce) != nonceSize {
  panic("invalid nonce size")
}

// Extract the actual ciphertext after the nonce
encryptedData := cipherText[nonceSize:]
if len(encryptedData) < 1 {
  panic("ciphertext is too short")
}

msg, err := gcm.Open(nil, nonce, encryptedData, nil)
if err != nil {
  panic("failed to decrypt message")
}


```
