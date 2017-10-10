package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"k8s.io/api/core/v1"
)

func sendMessageToSlackChannel(ev *v1.Event, url string) {
	var buffer bytes.Buffer

	msg := SlackMessage{
		InvolvedObjectKind: ev.InvolvedObject.Kind,
		InvolvedObjectName: ev.InvolvedObject.Name,
		Type:               ev.Type,
		Reason:             ev.Reason,
		Message:            ev.Message,
		Namesapce:          ev.Namespace,
	}

	buffer.WriteString(`{ "text": "`)
	buffer.WriteString(msg.toString())
	buffer.WriteString(`"}`)
	fmt.Println(buffer.String())
	http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(buffer.String()))
}

type SlackMessage struct {
	InvolvedObjectName string
	InvolvedObjectKind string
	Type               string
	Reason             string
	Message            string
	Namesapce          string
}

func (sm SlackMessage) toString() string {
	return fmt.Sprintf("Resource name: %s.\nResoucre kind: %s.\nEvent type: %s.\nEvent reason: %s.\nEvent message: %s.\nNamespace: %s", sm.InvolvedObjectName, sm.InvolvedObjectKind, sm.Type, sm.Reason, sm.Message, sm.Namesapce)
}
