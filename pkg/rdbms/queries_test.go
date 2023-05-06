package rdbms_test

import (
	"context"
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg"
	"github.com/Goboolean/shared-packages/pkg/rdbms"
	"github.com/joho/godotenv"
)



var (
	DB *rdbms.PSQL
	Queries *rdbms.Queries
)


func TestMain(m *testing.M) {
	
	if err := os.Chdir("../../"); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	DB = rdbms.NewDB(&pkg.Config{
		Host:     os.Getenv("PSQL_HOST"),
		Port:     os.Getenv("PSQL_PORT"),
		User:     os.Getenv("PSQL_USER"),
		Password: os.Getenv("PSQL_PASS"),
		Database: os.Getenv("PSQL_DATABASE"),
	})

	Queries = rdbms.New(DB)

	code := m.Run()

	os.Exit(code)
}



func TestCreateAccssInfo(t *testing.T) {
	if err := Queries.CreateAccessInfo(context.Background()); err != nil {
		t.Errorf("CreateAccessInfo() failed: %v", err)
	}
}