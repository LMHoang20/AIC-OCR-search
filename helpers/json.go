package helpers

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJSON(jsonFile *os.File) map[string]interface{} {
	byteValue, _ := io.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	return result
}
