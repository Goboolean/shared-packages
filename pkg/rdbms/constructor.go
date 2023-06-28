package rdbms

import (
	"database/sql"
	"fmt"

	"github.com/Goboolean/shared-packages/pkg/resolver"
)



type PSQL struct {
	*sql.DB
}

func NewDB(c *resolver.ConfigMap) *PSQL {

	user, err := c.GetStringKey("USER")
	if err != nil {
		panic(err)
	}

	password, err := c.GetStringKey("PASSWORD")
	if err != nil {
		panic(err)
	}

	host, err := c.GetStringKey("HOST")
	if err != nil {
		panic(err)
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		panic(err)
	}

	database, err := c.GetStringKey("DATABASE")
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	  host, port, user, password, database)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(db)
	}

	return &PSQL{DB: db}
}

func (p *PSQL) Close() error {
	return p.DB.Close()
}

