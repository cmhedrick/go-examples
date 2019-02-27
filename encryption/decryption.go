// example taken from https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func main() {
	// provide key used for encryption
	key := []byte("passphrasewhichneedstobe32bytes!")
	// file to decrypt
	ciphertext, err := ioutil.ReadFile("myfile.data")

	if err != nil {
		fmt.Println(err)
	}

	// build cipher used for encryption originally
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	// set up galois counter mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	// create nonce that matches that used with original gcm
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plaintext))
}
