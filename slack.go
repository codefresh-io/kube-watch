package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"k8s.io/api/core/v1"
)

func sendMessageToSlackChannel(ev *v1.Event, url string) {
	var buffer bytes.Buffer

	msg := slackMessage{
		InvolvedObjectKind: ev.InvolvedObject.Kind,
		InvolvedObjectName: ev.InvolvedObject.Name,
		Type:               ev.Type,
		Reason:             ev.Reason,
		Message:            ev.Message,
		Namesapce:          ev.Namespace,
	}

	payload := strings.NewReader("{\n\t\"text\": \"hisda\"\n}")

	buffer.WriteString(`{ "text": "`)
	buffer.WriteString(msg.toString())
	buffer.WriteString(`", "icon_emoji": ":watch:"}`)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "20b801ea-562c-4f18-8188-74d2c02dca31")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(buffer.String()))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res.Status)
}

type slackMessage struct {
	InvolvedObjectName string
	InvolvedObjectKind string
	Type               string
	Reason             string
	Message            string
	Namesapce          string
}

func (sm slackMessage) toString() string {
	return fmt.Sprintf("Resource name: %s.\nResoucre kind: %s.\nEvent type: %s.\nEvent reason: %s.\nEvent message: %s.\nNamespace: %s", sm.InvolvedObjectName, sm.InvolvedObjectKind, formatEventType(sm.Type), sm.Reason, sm.Message, sm.Namesapce)
}

func formatEventType(str string) string {
	if str == "Warning" {
		return fmt.Sprintf("*%s*", str)
	}
	return str
}
