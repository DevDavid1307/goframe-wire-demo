package payment

import (
	"github.com/google/wire"
	"goframe/internal/infra/data/ent"
	"goframe/internal/mall/account"
)

var ProviderSet = wire.NewSet(
	ProviderHandler,
	ProviderService,
	ProviderDao,

	// bind interface
	wire.Bind(new(Controller), new(*controller)),
	wire.Bind(new(Service), new(*logic)),
	wire.Bind(new(Dao), new(*dao)),

	// dep other module
	account.ProviderExternalSet,
)

func ProviderHandler(svc Service) *controller {
	return &controller{svc: svc}
}

func ProviderService(repo Dao, account account.Service) *logic {
	return &logic{
		repo:    repo,
		account: account,
	}
}

func ProviderDao(db *ent.Client) *dao {
	return &dao{db: db}
}
