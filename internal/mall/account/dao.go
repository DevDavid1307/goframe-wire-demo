package account

import (
	"goframe/internal/infra/data/ent"
)

type Dao interface {
	GetUser() string
	Greeting() string
}

type dao struct {
	db *ent.Client
}

func (d *dao) GetUser() string {
	return "david"
}

func (d *dao) Greeting() string {
	return "Hello World!"
}
