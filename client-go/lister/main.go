package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/sumanjeet/.kube/config", "absolute path to the kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s building config from flags", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s getting in cluster config", err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s creating clientset", err.Error())
	}
	// fmt.Println(clientset)

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s listing pods", err.Error())
	}

	fmt.Println("The pods in the default namespace are:")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	// fmt.Println(pods)

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s listing deployments", err.Error())
	}

	fmt.Println("The deployments in the default namespace are:")
	for _, deployment := range deployments.Items {
		fmt.Println(deployment.Name)
	}

	time.Sleep(100 * time.Second) // Sleep for 10 seconds before next iteration

}
