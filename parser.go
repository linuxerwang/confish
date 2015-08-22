package confish

import (
	"errors"
	"io"
	"io/ioutil"
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

func Parse(r io.Reader, confVar interface{}) error {
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
