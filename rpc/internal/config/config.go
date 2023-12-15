package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf   config.DatabaseConf
	CasbinConf     casbin.CasbinConf
	RedisConf      redis.RedisConf
	KubeConfigPath K8sConfigPath
}
type K8sConfigPath struct {
	Path string
}
