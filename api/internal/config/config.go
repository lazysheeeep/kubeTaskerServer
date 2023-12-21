package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	RedisConf    redis.RedisConf
	CoreRpc      zrpc.RpcClientConf
	JobRpc       zrpc.RpcClientConf
	McmsRpc      zrpc.RpcClientConf
	Captcha      captcha.Conf
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
	I18nConf     i18n.Conf
	ProjectConf  ProjectConf
	CROSConf     config.CROSConf
}

type ProjectConf struct {
	DefaultRoleId           uint64 `json:"defaultRoleId,omitempty"`
	DefaultDepartmentId     uint64 `json:"defaultDepartmentId,omitempty"`
	DefaultPositionId       uint64 `json:"defaultPositionId,omitempty"`
	EmailCaptchaExpiredTime int    `json:"emailCaptchaExpiredTime,omitempty"`
	SmsTemplateId           string `json:"smsTemplateId,omitempty"`
	SmsAppId                string `json:"smsAppId,omitempty"`
	SmsSignName             string `json:"smsSignName,omitempty"`
	RegisterVerify          string `json:"registerVerify,omitempty"`
	LoginVerify             string `json:"loginVerify,omitempty"`
	ResetVerify             string `json:"resetVerify,omitempty"`
	AllowInit               bool   `json:"allowInit,omitempty"`
}
