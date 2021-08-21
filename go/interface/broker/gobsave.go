package broker

import (
	"encoding/gob"
	"os"

	stg "../strategy"
)


// saveGob save as gob file
func (b *Binance) SaveGob(file string) error {
	f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
	gob.Register(&stg.KeepStrategy{})
	encode := gob.NewEncoder(f)
	if err := encode.Encode(b); err != nil {
		return err
	}
	defer f.Close()
	return nil
}

// ReadGob read gob file
func (b *Binance) ReadGob(file string) error {
	f, _ := os.Open(file)
	gob.Register(&stg.KeepStrategy{})
	decode := gob.NewDecoder(f)
	if err := decode.Decode(b); err != nil {
		return err
	}
	os.Remove(file)
	return nil
}
