package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/sumanjeet/.kube/config", "absolute path to the kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// fmt.Println(clientset)

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("The pods in the default namespace are:")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	// fmt.Println(pods)

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("The deployments in the default namespace are:")
	for _, deployment := range deployments.Items {
		fmt.Println(deployment.Name)
	}

}
