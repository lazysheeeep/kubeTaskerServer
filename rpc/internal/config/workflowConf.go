package config

import (
	"github.com/kubeTasker/kubeTasker/pkg/client/clientset/versioned"
	"github.com/kubeTasker/kubeTasker/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type WorkflowConf struct {
	RestConfig   *rest.Config
	ClientConfig clientcmd.ClientConfig
	WfClient     v1alpha1.WorkflowInterface
	Clientset    *kubernetes.Clientset
}

func InitWorkflowClient(wfconfig *WorkflowConf, ns ...string) {
	if wfconfig.WfClient != nil {
		return
	}
	InitKubeClient(wfconfig)
	var namespace string
	var err error
	if len(ns) > 0 {
		namespace = ns[0]
	} else {
		namespace, _, err = wfconfig.ClientConfig.Namespace()
		if err != nil {
			log.Fatal(err)
		}
	}
	wfcs := versioned.NewForConfigOrDie(wfconfig.RestConfig)
	wfconfig.WfClient = wfcs.KubetaskerV1alpha1().Workflows(namespace)
}
