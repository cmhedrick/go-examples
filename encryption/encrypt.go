// example taken from https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	text := []byte("My Super Secret Code Stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	// create a cipher with var key
	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	// gcm = Galois/Counter Mode, a mode used for symmetric encryption
	// https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte list/array the size of the nonce
	// https://en.wikipedia.org/wiki/Cryptographic_nonce
	nonce := make([]byte, gcm.NonceSize())
	// randomize the values withing nonce
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	// prints out completed array of bytes
	fmt.Println(gcm.Seal(nonce, nonce, text, nil))

	// the WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}
