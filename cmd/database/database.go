package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func InitDatabase() {

	dsn := "host=" +
		os.Getenv("POSTGRES_IP") + " user=" +
		os.Getenv("POSTGRES_USER") + " password=" +
		os.Getenv("POSTGRES_PASSWORD") +
		" dbname= " + os.Getenv("DB_NAME") + " port=" +
		os.Getenv("POSTGRES_PORT") + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	for db == nil {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		//os.Exit(1)
		panic("Unable to connect to postgres database")

	}

	/*

		sqlDb, err := db.DB()

		if err != nil {
			panic(err)
		}
		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDb,
		}), &gorm.Config{})
		DbConn = gormDB
	*/

	PostgresDB = db

}
