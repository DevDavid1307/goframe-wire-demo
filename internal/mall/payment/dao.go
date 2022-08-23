package payment

import (
	"goframe/internal/infra/data/ent"
)

type Dao interface {
	Create()
}

type dao struct {
	db *ent.Client
}

func (d *dao) Create() {
}
