package config

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

func InitKubeClientConf(wfconfig *WorkflowConf) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()

	overrides := clientcmd.ConfigOverrides{}

	wfconfig.ClientConfig = clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, &overrides, os.Stdin)
}

func InitKubeClient(wfconfig *WorkflowConf) *kubernetes.Clientset {
	if wfconfig.Clientset != nil {
		return wfconfig.Clientset
	}
	var err error
	wfconfig.RestConfig, err = wfconfig.ClientConfig.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	wfconfig.Clientset, err = kubernetes.NewForConfig(wfconfig.RestConfig)
	if err != nil {
		log.Fatal(err)
	}
	return wfconfig.Clientset
}
