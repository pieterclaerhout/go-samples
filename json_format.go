package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func formatJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes(), err
	}
	return data, nil
}

func main() {

	data := []byte(`{"key":"hello","msg":"world"}`)

	prettyJSON, err := formatJSON(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(prettyJSON))

}
