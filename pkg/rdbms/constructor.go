package rdbms

import (
	"database/sql"
	"fmt"

	"github.com/Goboolean/shared-packages/pkg/resolver"
)



type PSQL struct {
	*sql.DB
}

func NewDB(c *resolver.Config) *PSQL {

	if err := c.ShouldHostExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldUserExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPortExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldPWExist(); err != nil {
		panic(err)
	}

	if err := c.ShouldDBExist(); err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	  c.Host, c.Port, c.User, c.Password, c.Database)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(db)
	}

	return &PSQL{DB: db}
}

func (p *PSQL) Close() error {
	return p.DB.Close()
}

