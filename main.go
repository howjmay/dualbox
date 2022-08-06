package main

import (
	"dualbox/pkg/crypt"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	cpt := crypt.NewCrypter(crypt.CRYPTO_TYPE_GCM)
	plaintext := "freedom and liberty"
	c, n, err := cpt.Encrypt([]byte(plaintext))
	if err != nil {
		logrus.Fatal(err)
	}
	decrypttext, err := cpt.Decrypt(c, n)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println("plaintext  : ", plaintext)
	fmt.Println("decrypttext: ", string(decrypttext))
}
