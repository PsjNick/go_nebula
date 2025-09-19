package nebula

import (
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

type NebulaLogger struct {
	ctx context.Context
}

func (l NebulaLogger) Info(msg string) {
	glog.Info(l.ctx, msg)
}

func (l NebulaLogger) Warn(msg string) {
	glog.Warning(l.ctx, msg)
}

func (l NebulaLogger) Error(msg string) {
	glog.Error(l.ctx, msg)
}

func (l NebulaLogger) Fatal(msg string) {
	glog.Fatal(l.ctx, msg)
}
