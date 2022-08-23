package payment

import "goframe/internal/mall/account"

type logic struct {
	repo Dao

	// 依赖其它模块
	account account.Service
}

func (l *logic) Pay() string {
	return "payment from " + l.account.GetUser()
}
