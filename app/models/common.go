package models

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func BindJsonParams(i io.Reader, s interface{}) error {
	bytes, err := ioutil.ReadAll(i)
	if err != nil {
		return nil
	}

	if len(bytes) == 0 {
		return nil
	}

	return json.Unmarshal(bytes, s)
}
