package jsonl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/youpy/go-jsonl"
)

type foo struct {
	Bar string `json:"bar"`
}

func TestUnmarshal(t *testing.T) {
	var result []foo

	jsonlData := `{"bar":"aaa"}
{"bar":"bbb"}`

	err := jsonl.Unmarshal([]byte(jsonlData), &result)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[1].Bar, "bbb")
}
