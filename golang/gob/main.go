package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Hoge struct {
	F1 string
	F2 int64
}

type Fuga struct {
	F3 string
	F4 int
}

func main() {
	encoded := encode(Hoge{F1: "hoge", F2: 123})
	fmt.Printf("encoded: %d bytes\n", len(encoded))

	var hoge Hoge
	decode(&hoge, encoded)
	fmt.Printf("hoge: %+v\n", hoge)

	encoded = encode(Fuga{F3: "fuga", F4: 456})
	fmt.Printf("encoded: %d bytes\n", len(encoded))

	var fuga Fuga
	decode(&fuga, encoded)
	fmt.Printf("fuga: %+v\n", fuga)
}

func encode(item interface{}) []byte {
	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(item)
	return buf.Bytes()
}

func decode(dest interface{}, data []byte) error {
	buf := bytes.NewBuffer(data)
	return gob.NewDecoder(buf).Decode(dest)
}
