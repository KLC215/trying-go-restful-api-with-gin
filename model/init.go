package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func InitSelfDB() *gorm.DB {
	return connectDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.name"))
}

func InitDockerDB() *gorm.DB {
	return connectDB(
		viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.host"),
		viper.GetString("docker_db.name"))
}

func connectDB(username, password, host, name string) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username, password, host, name, true, "Local")

	db, err := gorm.Open("mysql", connStr)

	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	// Setting maximum connections can avoid mysql "too many connections" error
	//db.DB().SetMaxOpenConns(20000)

	// Reuse mysql connection in connection pool
	db.DB().SetMaxIdleConns(0)
}
