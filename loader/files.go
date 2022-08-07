package loader

import (
	"dualbox/utils"
	"io/ioutil"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const FilenameExtension = ".enc"

// encoding standard
// | file_name_length (4 bytes) | file_name | data ..... |
func OpenFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Fatal(err)
	}

	fileName := path.Base(filePath)
	fileNameLen := uint32(len(fileName))
	buf := utils.Uint32ToBytes(fileNameLen)
	buf = append(buf, []byte(fileName)...)
	buf = append(buf, data...)
	return buf, nil
}

func WriteFile(newName string, b []byte) error {
	nameLen := utils.BytesToUint32(b[:4])
	var name string
	if newName == "" {
		name = string(b[4 : 4+nameLen])
	} else {
		name = newName
	}
	data := b[4+nameLen:]

	err := os.WriteFile(name, data, 0644)
	if err != nil {
		logrus.Fatal(err)
	}
	return nil
}
