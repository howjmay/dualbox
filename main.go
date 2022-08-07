package main

import (
	"dualbox/pkg/crypt"
	"dualbox/pkg/loader"
	"dualbox/pkg/pack"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
}

func main() {
	cpt := crypt.NewCrypter(crypt.CRYPTO_TYPE_GCM_AES256)

	b0, err := loader.OpenFile("./testdata/testdata1.jpg")
	if err != nil {
		logrus.Fatal(err)
	}

	b1, err := loader.OpenFile("./testdata/testdata2.png")
	if err != nil {
		logrus.Fatal(err)
	}

	key0 := cpt.GenRandKey(crypt.CRYPTO_TYPE_GCM_AES256)
	key1 := cpt.GenRandKey(crypt.CRYPTO_TYPE_GCM_AES256)
	c0, n0, err := cpt.Encrypt(key0, nil, b0)
	if err != nil {
		logrus.Fatal(err)
	}
	c1, n1, err := cpt.Encrypt(key1, nil, b1)
	if err != nil {
		logrus.Fatal(err)
	}

	b, err := pack.Pack(cpt, key0, key1, n0, n1, c0, c1)
	if err != nil {
		logrus.Fatal(err)
	}

	cipherFilename := time.Now().Format(time.RFC3339) + loader.FilenameExtension
	err = loader.WriteFile(cipherFilename, b)
	if err != nil {
		logrus.Fatal(err)
	}

	decrypttext0, err := pack.Unpack(cpt, b, key0)
	if err != nil {
		logrus.Fatal(err)
	}

	err = loader.WriteFile("", decrypttext0)
	if err != nil {
		logrus.Fatal(err)
	}
}
