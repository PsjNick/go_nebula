package nebula

import (
	"fmt"
	"time"

	"github.com/PsjNick/go_nebula/config"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	nebula_go "github.com/vesoft-inc/nebula-go/v3"
)

var (
	nebulaConfig config.NebulaConfig
	err          error
)

var NebulaSessionPool *nebula_go.SessionPool

func InitNebula(config config.NebulaConfig) (err error) {
	nebulaConfig = config

	ctx := gctx.GetInitCtx()

	var serviceAddrs []nebula_go.HostAddress

	for _, host := range nebulaConfig.Hosts {
		hostAddress := nebula_go.HostAddress{Host: host.Host, Port: host.Port}
		serviceAddrs = append(serviceAddrs, hostAddress)
	}

	poolConfig, err := nebula_go.NewSessionPoolConf(
		nebulaConfig.Username,
		nebulaConfig.Password,
		serviceAddrs,
		nebulaConfig.SpaceName,
		[]nebula_go.SessionPoolConfOption{
			// 设置最大连接数
			nebula_go.WithMaxSize(nebulaConfig.Pool.MaxSize),
			// 设置最小连接数
			nebula_go.WithMinSize(nebulaConfig.Pool.MinSize),
			// 设置会话超时时间
			nebula_go.WithTimeOut(time.Duration(nebulaConfig.Pool.Timeout) * time.Second),
			// 设置空闲时间
			nebula_go.WithIdleTime(time.Duration(nebulaConfig.Pool.IdleTime) * time.Second),
		}...,
	)

	if err != nil {
		glog.Error(ctx, "failed to create session pool config", err.Error())
		return
	}

	// 创建 session pool
	NebulaSessionPool, err = nebula_go.NewSessionPool(*poolConfig, NebulaLogger{ctx: ctx})

	if err != nil {
		glog.Error(ctx, "failed to create session pool config", err.Error())
		return
	}

	// 检查如果 session pool is valid
	if NebulaSessionPool == nil {
		glog.Error(ctx, "session pool is nil")
		err = fmt.Errorf("session pool is nil")
		return
	}

	glog.Info(ctx, "nebula session pool init success")

	return
}
