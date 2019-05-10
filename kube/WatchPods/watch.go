package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// This program lists the pods in a cluster equivalent to
//
// kubectl get pods
//
func main() {
	var ns string
	flag.StringVar(&ns, "namespace", "", "namespace for listing pods")
	flag.Parse()
	// Bootstrap k8s configuration from local 	Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Println("Using kubeconfig file: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create an rest client not targeting specific API version
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}

	// print pods
	for i, pod := range pods.Items {
		fmt.Printf("[%d] Namespace: %s %s %s\n", i, pod.Namespace, pod.GetName(), pod.GetLabels())
	}

	watcher, err := clientset.CoreV1().Pods(ns).Watch(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error in creating watcher %s\n", err.Error())
	}
	ch := watcher.ResultChan()
	for event := range ch {
		pod, ok := event.Object.(*v1.Pod)
		if !ok {
			continue
		}

		fmt.Printf("Event Generated %v\n", event.Type)
		fmt.Printf("Event %s\n", event.Object.GetObjectKind())
		fmt.Printf("Pod name is %s\n", pod.Name)
	}
}
