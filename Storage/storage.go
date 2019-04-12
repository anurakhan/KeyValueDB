package storage

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
	"path/filepath"
)

//entity that controls writing and reading data from disk
type DataBaseDiskController struct {
	FilePath string
}

//Saves data in generic interface to disk in binary encoding.
func (cntr *DataBaseDiskController) SaveToDisk(data map[interface{}]interface{}) bool {
	filePath, _ := filepath.Abs(cntr.FilePath)
	file, err := os.Create(filePath + "/dbData")
	defer file.Close()

	if err != nil {
		return false
	}
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(data)
	resToSave := buf.Bytes()

	file.Write(resToSave)

	return true
}

//Reads data from disk and returns it.
func (cntr *DataBaseDiskController) ReadFromDisk() (map[interface{}]interface{}, bool) {
	filePath, _ := filepath.Abs(cntr.FilePath)
	dat, err := ioutil.ReadFile(filePath + "/dbData")

	if err != nil {
		return nil, false
	}

	reader := bytes.NewReader(dat)
	decoder := gob.NewDecoder(reader)
	var obj map[interface{}]interface{}
	decoder.Decode(&obj)

	return obj, true
}
