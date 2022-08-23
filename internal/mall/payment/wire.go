//go:build wireinject
// +build wireinject

package payment

import (
	"github.com/google/wire"
	"goframe/internal/infra/data/ent"
)

func InitPayment(db *ent.Client) *controller {
	panic(wire.Build(ProviderSet))
}
