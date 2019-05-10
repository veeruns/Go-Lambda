package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

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
		panic(err.Error())
	}
	//Create informer factory
	factory := informers.NewSharedInformerFactory(clientset, 0)
	//Pod Creating informer
	informer := factory.Core().V1().Pods().Informer()
	//	informer := factory.Apps().V1().Deployments().Informer()
	stopper := make(chan struct{})
	defer close(stopper)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			// "k8s.io/apimachinery/pkg/apis/meta/v1" provides an Object
			// interface that allows us to get metadata easily
			mObj := obj.(v1.Object)
			log.Printf("New Pod Added to Store: %v", mObj.GetName())
		},
		UpdateFunc: func(old interface{}, obj interface{}) {
			mObj := obj.(v1.Object)
			oObj := old.(v1.Object)
			log.Printf("Pod was updated in %s from %s (was %s)\n", mObj.GetNamespace(), mObj.GetName(), oObj.GetName())

		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			t := mObj.GetDeletionTimestamp()
			log.Printf("Pod was deleted %s %s\n", mObj.GetName(), t)
		},
	})

	go informer.Run(stopper)

	<-stopper
}
