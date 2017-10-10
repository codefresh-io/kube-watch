package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
)

func doPost(obj interface{}, url string) {
	mar, _ := json.Marshal(obj)
	jsonStr := []byte(string(mar))
	if url == "" {
		fmt.Println(string(mar))
	} else {
		http.Post(url, "Application/json", bytes.NewBuffer(jsonStr))
	}

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
