package v1

import "github.com/gogf/gf/v2/frame/g"

type PaymentIndexReq struct {
	g.Meta `path:"/index" tags:"Hello" method:"get" summary:"You first hello api"`
}

type PaymentIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
