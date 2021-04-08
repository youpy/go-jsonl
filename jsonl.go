package jsonl

import (
	"bytes"
	"encoding/json"
)

func Unmarshal(data []byte, v interface{}) error {
	lines := bytes.Split(data, []byte("\n"))
	joined := bytes.Join(lines, []byte(","))
	jsonData := bytes.Join([][]byte{[]byte("["), joined, []byte("]")}, []byte(""))
	err := json.Unmarshal(jsonData, v)
	if err != nil {
		return err
	}

	return nil
}
