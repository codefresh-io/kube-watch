package main

import (
	"fmt"
	"time"

	"github.com/urfave/cli"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/cache"
)

func watch(clientset *kubernetes.Clientset, context *cli.Context) {
	watchlist := cache.NewListWatchFromClient(clientset.Core().RESTClient(), "events", v1.NamespaceAll, fields.Everything())
	_, controller := cache.NewInformer(
		watchlist,
		&v1.Event{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				onAdd(obj, context)
			},
		},
	)
	stop := make(chan struct{})
	go controller.Run(stop)
	for {
		time.Sleep(time.Second)
	}
}

func onAdd(obj interface{}, context *cli.Context) {
	ev := obj.(*v1.Event)
	fmt.Println(ev.InvolvedObject.Kind)
	watchType := context.String("watch-type")
	if ev.Type == watchType || watchType == "ALL" {
		if context.IsSet("url") == true {
			doPost(obj, context.String("url"))
		}
		if context.IsSet("slack-channel-url") == true {
			sendMessageToSlackChannel(ev, context.String("slack-channel-url"))
		}
	}
}
