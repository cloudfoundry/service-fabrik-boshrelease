
//TODO copyright header


package main

import (
	"flag"
	"log"

	controllerlib "github.com/kubernetes-incubator/apiserver-builder/pkg/controller"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller"
)

var kubeconfig = flag.String("kubeconfig", "", "path to kubeconfig")

func main() {
	flag.Parse()
	config, err := controllerlib.GetConfig(*kubeconfig)
	if err != nil {
		log.Fatalf("Could not create Config for talking to the apiserver: %v", err)
	}

	controllers, _ := controller.GetAllControllers(config)
	controllerlib.StartControllerManager(controllers...)

	// Blockforever
	select {}
}
