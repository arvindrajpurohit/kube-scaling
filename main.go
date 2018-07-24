package main

import (
	"flag"
	"fmt"

	kube "github.com/arvindrajpurohit/kube-scaling/kube"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//var replicaset_name = flag.String(replicaset_name )
	var ip = flag.Int("flagname", 1234, "help message for flagname")

	config, err := clientcmd.BuildConfigFromFlags("", "/Users/admin/.kube/config")
	fmt.Println(config.ServerName)
	clientset, err := kubernetes.NewForConfig(config)
	fmt.Println(clientset)
	if err != nil {
		panic(err)
	}

	a := kube.Scalereplicaset("frontend", "default", &ip)
	fmt.Println(a)
	scaleservice1, err := clientset.ExtensionsV1beta1().ReplicaSets(a.Namespace).UpdateScale(a.Name, a)
	fmt.Println(scaleservice1.Name+" Replicaset scale to ", scaleservice1.Spec.Replicas)

}
