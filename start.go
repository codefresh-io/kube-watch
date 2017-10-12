package main

import (
	"bytes"
	"fmt"

	"github.com/urfave/cli"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func dryRun(c *cli.Context) {
	fmt.Println("Running on current-context from", c.String("kube-config"))
	var kubeconfig string
	var buffer bytes.Buffer
	buffer.WriteString(c.String("kube-config"))
	kubeconfig = buffer.String()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("kubeconfig ", kubeconfig)
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err == nil {
		watch(clientset, c)
	} else {
		panic(err)
	}
}
