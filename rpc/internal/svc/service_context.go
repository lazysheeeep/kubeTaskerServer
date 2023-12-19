package svc

import (
	"github.com/kubeTasker/kubeTaskerServer/rpc/ent"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"k8s.io/client-go/kubernetes"

	_ "github.com/kubeTasker/kubeTaskerServer/rpc/ent/runtime"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  *redis.Redis
	K8s    *kubernetes.Clientset
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	k8s, err := K8sInit()
	if err != nil {
		return nil, err
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis.MustNewRedis(c.RedisConf),
		K8s:    k8s,
	}, nil
}
