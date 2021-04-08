package jsonl

import (
	"bytes"
	"encoding/json"
)

func Unmarshal(data []byte, v interface{}) error {
	lines := bytes.Split(data, []byte("\n"))
	filtered := [][]byte{}
	for _, line := range lines {
		if len(line) > 0 {
			filtered = append(filtered, line)
		}
	}
	joined := bytes.Join(filtered, []byte(","))
	jsonData := bytes.Join([][]byte{[]byte("["), joined, []byte("]")}, []byte(""))
	err := json.Unmarshal(jsonData, v)
	if err != nil {
		return err
	}

	return nil
}
