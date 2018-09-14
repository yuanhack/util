package util // import "fixup.cc/go/util"

import (
	"bytes"
	"encoding/gob"
)

// gob
// 将结构体编码成 []byte
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// gob
// 将 []byte 解码成 interface 指定的结构体
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
