package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// This program lists the pods in a cluster equivalent to
//
// kubectl get pods
//

type K8sResult struct {
	Context string
	Output  []string
}

func worker(newctx string, result chan K8sResult) {
	var emit K8sResult
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	//log.Println("Using kubeconfig file: ", kubeconfig)
	configOverrides := clientcmd.ConfigOverrides{}
	configOverrides.CurrentContext = newctx
	cfg, _ := clientcmd.LoadFromFile(kubeconfig)
	fmt.Printf("Contexts in kubeconfig worker %s\n", newctx)
	//config, _ := clientcmd.NewNonInteractiveClientConfig(&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig}, &cfgoverrides, "").ClientConfig()
	config, err := clientcmd.NewNonInteractiveClientConfig(*cfg, newctx, &configOverrides, nil).ClientConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("From Config %s\n", config.CertFile)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	var op []string
	for _, v := range nodes.Items {
		op = append(op, v.GetName())
	}
	emit.Context = newctx
	emit.Output = op

	result <- emit
	return
}
func main() {
	var ns string
	var result = make(chan K8sResult)
	flag.StringVar(&ns, "namespace", "", "namespace for listing pods")
	flag.Parse()
	//results := make(chan string)
	// Bootstrap k8s configuration from local 	Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	cfg, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	for key := range cfg.Contexts {
		configOverrides := clientcmd.ConfigOverrides{}
		configOverrides.CurrentContext = key

		go worker(key, result)

	}
	for i := range result {
		fmt.Printf("%s %v\n", i.Context, i.Output)
		//fmt.Printf("Return from chan is %v\n", m)
	}
	close(result)

}
