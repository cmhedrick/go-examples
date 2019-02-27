package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

func readFile(inFile string) string {
	data, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println("[X] Failed to read")
	}

	return string(data)
}

func writeFile(newFile string, fileData string) bool {
	file, err := os.Create(newFile)
	if err != nil {
		return false
	}
	defer file.Close()

	file.WriteString(fileData)
	return true
}

func encryptFile(fileName string, passPhrase string) {
	text := []byte(readFile(fileName))
	key := []byte(passPhrase)

	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("encrypted.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func decryptFile(fileName string, passPhrase string) {
	key := []byte(passPhrase)
	ciphertext, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

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

func main() {
	// init cli app
	app := cli.NewApp()
	// set name and usage
	app.Name = "File Encrypt/Decrypt CLI"
	app.Usage = "Encrypt/Decrypt files"

	// define flags
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "test.txt",
		},
		cli.StringFlag{
			Name:  "passphrase, p",
			Value: "passphrasewhichneedstobe32bytes!",
		},
	}

	// create commands in dictionary like objects
	app.Commands = []cli.Command{
		{
			Name:  "encrypt",
			Usage: "encrypt files",
			Flags: myFlags,
			// action will be used to run functions
			Action: func(c *cli.Context) error {
				fileName := c.String("file")
				passPhrase := c.String("passphrase")
				encryptFile(fileName, passPhrase)
				return nil
			},
		},
		{
			Name:  "decrypt",
			Usage: "decrypt files",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fileName := c.String("file")
				passPhrase := c.String("passphrase")
				decryptFile(fileName, passPhrase)
				return nil
			},
		},
	}

	// start our application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
