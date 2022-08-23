package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/infra/data"
	"goframe/internal/mall/account"
	"goframe/internal/mall/payment"
)

func main() {
	s := g.Server()

	// infra
	db := data.NewEntClient()

	// account
	s.Group("/account", func(group *ghttp.RouterGroup) {
		accountModule := account.InitAccount(db)

		group.Bind(accountModule)
	})

	// payment
	s.Group("/payment", func(group *ghttp.RouterGroup) {
		paymentModule := payment.InitPayment(db)

		group.Bind(paymentModule)
	})

	s.Run()
}
