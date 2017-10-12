package main

import (
	"fmt"

	"github.com/urfave/cli"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func installInCluster(c *cli.Context) {
	clientset := prepareLocalKubernetesClientOrPanic(c)
	deploymentsClient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	deployment := &appsv1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "kube-watch",
		},
		Spec: appsv1beta1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "kube-watch",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            "kube-watch",
							ImagePullPolicy: "Always",
							Image:           "olsynt/kubewatch:command-start",
							Args:            prepareContainerArgs(c),
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

}
func int32Ptr(i int32) *int32 { return &i }

func prepareContainerArgs(c *cli.Context) []string {
	var strs []string
	strs = append(strs, "run")
	strs = append(strs, "--in-cluster")
	if c.IsSet("url") {
		strs = append(strs, "--url")
		strs = append(strs, c.String("url"))
	}

	if c.IsSet("kube-config") {
		strs = append(strs, "--kube-config")
		strs = append(strs, c.String("kube-config"))
	}

	if c.IsSet("slack-channel-url") {
		strs = append(strs, "--slack-channel-url")
		strs = append(strs, c.String("slack-channel-url"))
	}

	if c.IsSet("watch-type") {
		strs = append(strs, "--watch-type")
		strs = append(strs, c.String("watch-type"))
	}

	if c.IsSet("watch-kind") {
		strs = append(strs, "--watch-kind")
		strs = append(strs, c.String("watch-kind"))
	}

	return strs
}
