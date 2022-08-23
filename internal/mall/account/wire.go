//go:build wireinject
// +build wireinject

package account

import (
	"github.com/google/wire"
	"goframe/internal/infra/data/ent"
)

func InitAccount(db *ent.Client) *controller {
	panic(wire.Build(ProviderSet))
}
