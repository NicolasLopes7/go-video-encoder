package database

import (
	"encoder/src/domain"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Database struct {
	Db             *gorm.DB
	Dsn            string
	DsnTest        string
	DbType         string
	DbTypeTest     string
	Debug          bool
	AutorMigrateDb bool
	Env            string
}

func NewDb() *Database {
	return &Database{}
}

func (db *Database) Connect() (*gorm.DB, error) {
	var err error

	if db.Env != "test" {
		db.Db, err = gorm.Open(db.DbType, db.Dsn)
	} else {
		db.Db, err = gorm.Open(db.DbTypeTest, db.DsnTest)
	}

	if err != nil {
		return nil, err
	}

	if db.Debug {
		db.Db.LogMode(true)
	}

	if db.AutorMigrateDb {
		db.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
		db.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos(id)", "CASCADE", "CASCADE")
	}

	return db.Db, nil
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutorMigrateDb = true
	dbInstance.Debug = true

	conn, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return conn
}
