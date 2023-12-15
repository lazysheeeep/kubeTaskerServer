package svc

import (
	"github.com/kubeTasker/kubeTaskerServer/rpc/ent"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	_ "github.com/kubeTasker/kubeTaskerServer/rpc/ent/runtime"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  *redis.Redis
	K8s    *kubernetes.Clientset
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	K8sRestConfig, err := clientcmd.BuildConfigFromFlags("", c.KubeConfigPath.Path)
	if err != nil {
		panic("Failed to obtain the K8s configuration:" + err.Error())
		return nil
	} else {
		logx.Info("K8s configuration successfully obtained!")
	}
	k8sClient, err := kubernetes.NewForConfig(K8sRestConfig)
	if err != nil {
		logx.Error("failed to connect to the k8s, please check the k8s configuration.")
		return nil
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis.MustNewRedis(c.RedisConf),
		K8s:    k8sClient,
	}
}
