package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AccountGreetingReq struct {
	g.Meta `path:"/greeting" tags:"Hello" method:"get" summary:"You first hello api"`
}

type AccountGreetingRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
