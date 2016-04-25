package main

import (
	"log"

	"k8s.io/kubernetes/pkg/api"
	kcl "k8s.io/kubernetes/pkg/client/unversioned"
)

func main() {
	cl, err := kcl.NewInCluster()
	if err != nil {
		log.Fatalf("Error creating new k8s client (%s)", err)
	}
	ns := cl.Namespaces()
	watch, err := ns.Watch(api.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing namespaces (%s)", err)
	}
	ch := watch.ResultChan()
	for evt := range ch {
		log.Printf("%s %s", evt.Type, evt.GetObjectKind())
	}
	log.Fatalf("Events channel was closed")
}
