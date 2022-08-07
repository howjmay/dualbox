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
func OpenPlainFile(filePath string) ([]byte, error) {
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
	var name string
	var data []byte
	if newName == "" {
		nameLen := utils.BytesToUint32(b[:4])
		name = string(b[4 : 4+nameLen])
		data = b[4+nameLen:]
	} else {
		name = newName
		data = b
	}

	err := os.WriteFile(name, data, 0644)
	if err != nil {
		logrus.Fatal(err)
	}
	return nil
}
