# go-jsonl

## Usage

```go
import "github.com/youpy/go-jsonl"

type foo struct {
    Bar string `json:"bar"`
}

func main() {
    var result []foo

    jsonlData := `{"bar":"aaa"}
{"bar":"bbb"}`

    err := jsonl.Unmarshal([]byte(jsonlData), &result)
}
```
