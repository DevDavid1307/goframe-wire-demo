package payment

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "goframe/api/v1"
)

type Controller interface {
	Pay(ctx context.Context, req *v1.PaymentIndexReq) (res *v1.PaymentIndexRes, err error)
}

type controller struct {
	svc Service
}

func (c *controller) Pay(ctx context.Context, req *v1.PaymentIndexReq) (res *v1.PaymentIndexRes, err error) {
	_ = g.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
		"msg": c.svc.Pay(),
	})
	return
}
