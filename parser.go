package confish

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ParseFile(confFile string, confVar interface{}) error {
	f, err := os.Open(confFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return Parse(f, confVar)
}

func Parse(r io.Reader, confVar interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("%v", r)
			err = errors.New(msg)
			log.Println("Error, ", err)
		}
	}()

	content, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	ret := cfgParse(newCfgLex([]byte(content), confVar))
	if ret != 0 {
		return errors.New("confish parse failed.")
	}
	return nil
}
