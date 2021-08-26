package utils

import (
	"encoding/gob"
	"os"
)

// Dat struct store date create from program
// TODO: other program settings
type Dat struct {
	Password    string
	StoragePath []string
}

// NewDat create a new dat
func NewDat() *Dat {
	return &Dat{}
}

// AddStorage add path of storage
func (d *Dat) AddStorage(path string) {
	d.StoragePath = append(d.StoragePath, path)
}

// SaveEncode encode Dat and save to file
func (d *Dat) SaveEncode(name string) error {
	// open file
	f, err := os.OpenFile(name, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
	defer f.Close()
	if err != nil {
		return Errs("Open Encode Dat Error: ", err)
	}

	// encode
	encode := gob.NewEncoder(f)
	if err := encode.Encode(d); err != nil {
		panic(Errs("Encode Dat Error: ", err))
	}
	return nil
}

// ReadDecode read Dat and decode
func (d *Dat) ReadDecode(name string) error {
	// check file exist
	// if not exist, return
	isFile, err := IsFile(name)
	if err != nil {
		return Errs("Open Decode Dat Error: ", err)
	} 
	if !isFile {
		return nil
	}

	// open && decode 
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return Errs("Decode Dat Error: ", err)
	}
	decode := gob.NewDecoder(f)
	if err := decode.Decode(d); err != nil {
		panic(Errs("Decode Dat Error: ", err))
	}

	return nil
}
