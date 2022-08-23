package account

import (
	"github.com/google/wire"
	"goframe/internal/infra/data/ent"
)

var ProviderExternalSet = wire.NewSet(
	ProviderService,
	ProviderDao,

	// bind interface
	wire.Bind(new(Service), new(*logic)),
	wire.Bind(new(Dao), new(*dao)),
)

var ProviderSet = wire.NewSet(
	ProviderHandler,
	ProviderService,
	ProviderDao,

	// bind interface
	wire.Bind(new(Controller), new(*controller)),
	wire.Bind(new(Service), new(*logic)),
	wire.Bind(new(Dao), new(*dao)),
)

func ProviderHandler(svc Service) *controller {
	return &controller{svc: svc}
}

func ProviderService(repo Dao) *logic {
	return &logic{
		repo: repo,
	}
}

func ProviderDao(db *ent.Client) *dao {
	return &dao{db: db}
}
