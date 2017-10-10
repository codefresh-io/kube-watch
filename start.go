package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func onStart(c *cli.Context) {
	dryRun := c.Bool("dry-run")
	if dryRun {
		runDryRun(c)
	}
}

func runDryRun(c *cli.Context) {
	fmt.Println("Running in dry-run mode")
	fmt.Println("Fetching context from config %s", c.String("kube-config"))
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	a := c.StringSlice("watch-on")
	if len(a) > 0 {
		fmt.Println("Watching only on", a)
	} else {
		fmt.Println("Watching on all resources")
	}
	watch(clientset)
}
