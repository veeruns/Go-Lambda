package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"reflect"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
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
	Ioutput interface{}
}

type job struct {
	id        int
	context   string
	clientset *kubernetes.Clientset
	command   string
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
		var errflag bool
		errflag = false

		nodes, err := j.clientset.CoreV1().Nodes().List(metav1.ListOptions{})

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
			emit.Ioutput = j.clientset.CoreV1().Nodes()

			results <- emit
			fmt.Println("Done sending result from  worker", id, " and ", j.context)
		} else {
			var op []string
			emit.id = id
			emit.context = j.context
			emit.Output = op
			emit.Ioutput = nodes
			fmt.Printf("Empty result because %v\n", errflag)
			results <- emit
		}
		wg.Done()
	}
}
func main() {

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	cfg, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	var numcontexts int
	numcontexts = len(cfg.Contexts)
	fmt.Printf("Number of contexts is %d\n", numcontexts)
	var csets = make([]*kubernetes.Clientset, numcontexts+1)
	configOverrides := clientcmd.ConfigOverrides{}
	var v int
	var ctxarray = make([]string, numcontexts)
	for j := range cfg.Contexts {

		configOverrides.CurrentContext = j
		config, err := clientcmd.NewNonInteractiveClientConfig(*cfg, j, &configOverrides, nil).ClientConfig()
		if err != nil {
			log.Fatal(err)
		}

		clientset, err := kubernetes.NewForConfig(config)
		_, errn := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
		if errn != nil {
			fmt.Printf("There was error for a cluster %s\n", j)
		} else {
			v++
			csets[v] = clientset
			ctxarray[v] = j
			fmt.Printf("added a client set for context %s\n", j)
		}
	}
	var workingcontexts int
	workingcontexts = v
	fmt.Printf("Number of Working contexts are %d %d\n ", v, workingcontexts)

	timetick := time.Tick(3 * time.Second)
	for now := range timetick {
		fmt.Printf("Running for time %v\n", now)
		jobs := make(chan job, workingcontexts)
		result := make(chan result, workingcontexts)
		jerrors := make(chan joberror, workingcontexts)
		for w := 1; w <= workingcontexts; w++ {
			go worker(w, &wg, jobs, result, jerrors)
		}
		var jid int
		for k := 1; k <= workingcontexts; k++ {
			jid++
			var l job
			l.id = k
			l.context = ctxarray[k]
			l.clientset = csets[k]
			l.command = "node"
			jobs <- l
			wg.Add(1)
			fmt.Println("Adding job to worker ", jid)
		}
		close(jobs)
		wg.Wait()
		close(result)
		fmt.Println("Size of channel is ", len(result))
		for i := range result {
			fmt.Printf("Result %d [%s] %d \n", i.id, i.context, len(i.Output))
			//fmt.Printf("Type is %v\n", reflect.TypeOf(i.Ioutput))
			switch v := i.Ioutput.(type) {
			case *corev1.NodeList:
				fmt.Printf("Type is %v\n", reflect.TypeOf(i.Ioutput))
				//nodes, _ := v.Items
				fmt.Printf("Number of items is %d\n", len(v.Items))
			case v1.NodeInterface:
				//fmt.Printf("W Type is %v %v\n", reflect.TypeOf(i.Ioutput), v)
				n, _ := v.List(metav1.ListOptions{})
				for _, l := range n.Items {
					fmt.Printf("Node Status %v\n", l.Status.Conditions)
					fmt.Printf("Nodename %s, node status\n", l.GetName())
				}
			default:
				fmt.Printf("Type is %v %v\n", reflect.TypeOf(i.Ioutput), v)
				fmt.Printf("Random stuff %v\n", reflect.ValueOf(v).Type)

			}

		}
		close(jerrors)

	}

}
