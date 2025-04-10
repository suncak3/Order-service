package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

var once sync.Once

var database *sqlx.DB

func GetConnection() *sqlx.DB {
	once.Do(connectionDB)
	return database
}

func connectionDB() {
	connection, driver := getDbConfig()
	db, err := sqlx.Connect(driver, connection)

	if err != nil {
		panic("Failed db connection " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("DB is not working")
	}

	database = db
}

func getDbConfig() (string, string) {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Failed to load config of database!")
	}

	connectionURL := os.Getenv("connectionURL")
	driver := os.Getenv("driver")

	return connectionURL, driver
}

func CloseConnection() {
	err := database.Close()

	if err != nil {
		panic("Close connection failed")
	}

}
