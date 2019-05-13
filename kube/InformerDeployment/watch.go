package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/api/apps/v1"
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
	informer := factory.Apps().V1().Deployments().Informer()
	informer = factory.Core().V1().Pods().Informer()
	//	informer := factory.Apps().V1beta2().Deployments().Informer()
	//	informer := factory.Apps().V1().Deployments().Informer()
	stopper := make(chan struct{})
	defer close(stopper)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			// "k8s.io/apimachinery/pkg/apis/meta/v1" provides an Object
			// interface that allows us to get metadata easily
			mObj := obj.(*v1.Deployment)
			labels := mObj.GetObjectMeta().GetLabels()
			var op string
			for k, v := range labels {
				op = op + k + "=" + v + ","

			}
			log.Printf("New Deployment Added to Store: %s (labels=%s)", mObj.GetObjectMeta().GetName(), op)
		},
		UpdateFunc: func(old interface{}, obj interface{}) {
			mObj := obj.(*v1.Deployment)
			oObj := old.(*v1.Deployment)
			labels := mObj.GetObjectMeta().GetLabels()
			var op string
			for k, v := range labels {
				op = op + k + "=" + v + ","

			}
			log.Printf("Deployment was updated in %s from %s (labels = %s) (was %s)\n", mObj.GetNamespace(), mObj.GetName(), op, oObj.GetName())

		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(*v1.Deployment)
			labels := mObj.GetObjectMeta().GetLabels()
			var op string
			for k, v := range labels {
				op = op + k + "=" + v + ","

			}

			log.Printf("Pod was deleted %s with labels %v\n", mObj.GetName(), op)
		},
	})

	go informer.Run(stopper)

	<-stopper
}
