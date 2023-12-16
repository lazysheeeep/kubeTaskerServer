package svc

import (
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sConfig struct {
	restConfig   *rest.Config
	clientConfig clientcmd.ClientConfig
	clientset    *kubernetes.Clientset
}

func K8sInit() *kubernetes.Clientset {
	var k8sConfig K8sConfig
	initKubeClientConfig(&k8sConfig)
	initKubeClient(&k8sConfig)
	return k8sConfig.clientset
}

func initKubeClientConfig(k8sConfig *K8sConfig) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()

	overrides := clientcmd.ConfigOverrides{}

	k8sConfig.clientConfig = clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, &overrides, os.Stdin)
}

func initKubeClient(k8sConfig *K8sConfig) *kubernetes.Clientset {
	if k8sConfig.clientset != nil {
		return k8sConfig.clientset
	}
	var err error
	k8sConfig.restConfig, err = k8sConfig.clientConfig.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	k8sConfig.clientset, err = kubernetes.NewForConfig(k8sConfig.restConfig)
	if err != nil {
		log.Fatal(err)
	}
	return k8sConfig.clientset
}
