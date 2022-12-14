// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package payment

import (
	"goframe/internal/infra/data/ent"
	"goframe/internal/mall/account"
)

// Injectors from wire.go:

func InitPayment(db *ent.Client) *controller {
	paymentDao := ProviderDao(db)
	accountDao := account.ProviderDao(db)
	accountLogic := account.ProviderService(accountDao)
	paymentLogic := ProviderService(paymentDao, accountLogic)
	paymentController := ProviderHandler(paymentLogic)
	return paymentController
}
