package main

import (
	"bytes"
	"fmt"

	"github.com/urfave/cli"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func prepareLocalKubernetesClientOrPanic(c *cli.Context) *kubernetes.Clientset {
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
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func prepareInClusterKubernetesClientOrPanic(c *cli.Context) *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
