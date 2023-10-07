package pg

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hamdiBouhani/GopherNet-golang/storage/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	Db *gorm.DB
}

func NewDBConn() *DBConn {
	return &DBConn{}
}

func (svc *DBConn) CreateConnection() error {
	log.Println("Using Postgres Database")
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	ssl := os.Getenv("DB_SSL")
	if ssl == "" {
		ssl = "disable"
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", os.Getenv("DB_URL"), port, os.Getenv("DB_USER"), os.Getenv("DB_DATABASE"), ssl, os.Getenv("DB_PASS"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error updating burrow: %v\n", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(5 * time.Second)
	svc.Db = db
	return nil
}

func (svc *DBConn) Close() error {
	db, err := svc.Db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (svc *DBConn) Migrate() error {
	return svc.Db.AutoMigrate(
		&model.Burrow{},
	)
}

func (svc *DBConn) Drop() error {
	return svc.Db.Migrator().DropTable(
		&model.Burrow{},
	)
}
