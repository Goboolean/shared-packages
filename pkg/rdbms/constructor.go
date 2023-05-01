package rdbms

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)



var (
	PSQL_HOST     = os.Getenv("PSQL_HOST")
	PSQL_PORT     = os.Getenv("PSQL_PORT")
	PSQL_USER     = os.Getenv("PSQL_USER")
	PSQL_PASS     = os.Getenv("PSQL_PASS")
	PSQL_DATABASE = os.Getenv("PSQL_DATABASE")
)

func init() {

	if _, exist := os.LookupEnv("PSQL_HOST"); !exist {
		log.Fatalf("error: %s enveironment variable required", "PSQL_HOST")
	}

	if _, exist := os.LookupEnv("PSQL_PORT"); !exist {
		log.Fatalf("error: %s enveironment variable required", "PSQL_PORT")
	}

	if _, exist := os.LookupEnv("PSQL_USER"); !exist {
		log.Fatalf("error: %s enveironment variable required", "PSQL_USER")
	}

	if _, exist := os.LookupEnv("PSQL_PASS"); !exist {
		log.Fatalf("error: %s enveironment variable required", "PSQL_PASS")
	}

	if _, exist := os.LookupEnv("PSQL_DATABASE"); !exist {
		log.Fatalf("error: %s enveironment variable required", "PSQL_DATABASE")
	}
}



type PSQL struct {
	*sql.DB
}


func NewInstance() *PSQL {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	PSQL_HOST, PSQL_PORT, PSQL_USER, PSQL_PASS, PSQL_DATABASE)
	
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(db)
	}

	return &PSQL{DB: db}
}


func (p *PSQL) Close() error {
	return p.DB.Close()
}



func SetTx(tx *Transaction) *Queries {
	return New(NewInstance()).WithTx(tx.tx)
}