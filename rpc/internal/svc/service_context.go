package svc

import (
	"github.com/kubeTasker/kubeTaskerServer/rpc/ent"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/config"
	"k8s.io/client-go/kubernetes"

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

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis.MustNewRedis(c.RedisConf),
		K8s:    K8sInit(),
	}
}
