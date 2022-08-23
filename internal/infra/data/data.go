package data

import (
	"log"

	"goframe/internal/infra/data/ent"

	_ "github.com/go-sql-driver/mysql"
)

func NewEntClient() *ent.Client {
	client, err := ent.Open(
		"mysql",
		"david:123456@tcp(127.0.0.1:3306)/test?parseTime=True",
	)
	if err != nil {
		log.Fatalln("...")
	}

	return client
}
