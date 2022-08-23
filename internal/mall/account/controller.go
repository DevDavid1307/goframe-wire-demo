package account

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "goframe/api/v1"
)

type Controller interface {
	Greeting(ctx context.Context, req *v1.AccountGreetingReq) (res *v1.AccountGreetingRes, err error)
}

type controller struct {
	svc Service
}

func (c *controller) Greeting(ctx context.Context, req *v1.AccountGreetingReq) (res *v1.AccountGreetingRes, err error) {
	_ = g.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
		"msg": c.svc.Greet(),
	})
	return nil, nil
}
