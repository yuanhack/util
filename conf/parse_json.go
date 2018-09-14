package conf // import "fixup.cc/go/util/conf"

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ParseJson(filename string, v interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}
