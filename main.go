package main

import (
	kube "github.com/arvindrajpurohit/kube-scaling/kube"
)

func main() {
	a := kube.scalereplicaset("frontend", "default", 3)

}
