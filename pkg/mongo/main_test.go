package mongo_test

import (
	"os"
	"testing"

	"github.com/Goboolean/shared-packages/pkg/mongo"
	"github.com/Goboolean/shared-packages/pkg/resolver"
	"github.com/joho/godotenv"
)

var db *mongo.DB



func TestMain(m *testing.M) {

	if err := os.Chdir("../../"); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	SetupMongo()
	code := m.Run()
	TeardownMongo()

	os.Exit(code)
}


func SetupMongo() {
	c := resolver.Config{
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		Password: os.Getenv("MONGO_PASS"),
		Database: os.Getenv("MONGO_DATABASE"),
	}
	db = mongo.NewDB(&c)
	queries = mongo.New(instance)
}

func TeardownMongo() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func Test_Mongo(t *testing.T) {
	if err := db.Ping(); err != nil {
		t.Errorf("Ping() failed, %v", err)
	}
}