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

type result struct {
	id      int
	context string
	Output  []string
}

type job struct {
	id      int
	context string
}

type joberror struct {
	id  int
	err error
}

type k8sResult struct {
	Context string
	Output  []string
}

func worker(id int, wg *sync.WaitGroup, jobs <-chan job, results chan<- result, errors chan<- joberror) {
	for j := range jobs {
		fmt.Println("From worker ", j.id, "processing job ", j.context)

		var emit result
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		//log.Println("Using kubeconfig file: ", kubeconfig)
		configOverrides := clientcmd.ConfigOverrides{}
		configOverrides.CurrentContext = j.context
		cfg, _ := clientcmd.LoadFromFile(kubeconfig)
		fmt.Printf("Contexts in kubeconfig worker %s\n", j.context)
		//config, _ := clientcmd.NewNonInteractiveClientConfig(&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig}, &cfgoverrides, "").ClientConfig()
		config, err := clientcmd.NewNonInteractiveClientConfig(*cfg, j.context, &configOverrides, nil).ClientConfig()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("From Config %s\n", config.CertFile)
		clientset, err := kubernetes.NewForConfig(config)
		var errflag bool
		errflag = false
		if err != nil {
			var er joberror
			er.id = id
			er.err = err
			errors <- er
			errflag = true

		}

		nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})

		if err != nil {
			var er joberror
			er.id = id
			er.err = err
			errors <- er
			errflag = true

		}

		if !errflag {
			var op []string
			for _, v := range nodes.Items {
				op = append(op, v.GetName())
			}
			emit.id = id
			emit.context = j.context
			emit.Output = op

			results <- emit
			fmt.Println("Done sending result from  worker", id, " and ", j.context)
		}
		wg.Done()
	}
}
func main() {
	var ns string
	//	var result = make(chan K8sResult)

	flag.StringVar(&ns, "namespace", "", "namespace for listing pods")
	flag.Parse()
	//results := make(chan string)
	// Bootstrap k8s configuration from local 	Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	cfg, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	var numcontexts int
	numcontexts = len(cfg.Contexts)
	fmt.Printf("Number of contexts is %d\n", numcontexts)
	jobs := make(chan job, numcontexts)
	result := make(chan result, numcontexts)
	jerrors := make(chan joberror, numcontexts)
	for w := 1; w <= numcontexts; w++ {
		go worker(w, &wg, jobs, result, jerrors)
	}
	var jid int
	for key := range cfg.Contexts {
		jid++
		var l job
		l.id = jid
		l.context = key
		jobs <- l
		wg.Add(1)
		fmt.Println("Adding job to worker ", jid)
	}
	close(jobs)
	wg.Wait()
	close(result)
	fmt.Println("Size of channel is ", len(result))
	for i := range result {
		fmt.Printf("Result %d \n", i.id)
		//fmt.Printf("Return from chan is %v\n", m)
	}

	/*	select {
		case err := <-jerrors:
			fmt.Println("Error ", err.err.Error(), "Job id ", err.id)
		default:
		}*/

}
