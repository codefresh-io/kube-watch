package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"strings"
)

func doPost(obj interface{}, url string) {
	mar, _ := json.Marshal(obj)
	var buffer bytes.Buffer
	buffer.WriteString(`{ "body":`)
	buffer.WriteString(string(mar))
	buffer.WriteString(`}`)
	http.Post(url, "Application/json", strings.NewReader(buffer.String()))
}

func GetBytes(key interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}
