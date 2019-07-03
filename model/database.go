package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dialectName string
var dialectArgs string

func AutoMigrateDatabase() {
	dialectName = os.Getenv("APP_DATABASE_DRIVER")
	if "" == dialectName {
		panic("APP_DATABASE_DRIVER not found")
	}
	mysql_addr := os.Getenv("APP_MYSQL_ADDR")
	mysql_user := os.Getenv("APP_MYSQL_USER")
	mysql_passwd := os.Getenv("APP_MYSQL_PASSWORD")
	mysql_database := os.Getenv("APP_MYSQL_DATABASE")
	sqlite_filepath := os.Getenv("APP_SQLITE_FILEPATH")

	if "mysql" == dialectName {
		if "" == mysql_user {
			panic("APP_MYSQL_USER not found")
		}
		if "" == mysql_passwd {
			panic("APP_MYSQL_PASSWORD not found")
		}
		if "" == mysql_database {
			panic("APP_MYSQL_DATABASE not found")
		}
		if "" == mysql_addr {
			panic("APP_MYSQL_ADDR not found")
		}
	} else if "sqlite" == dialectName {
		dialectName = "sqlite3"
		if "" == sqlite_filepath {
			panic("APP_SQLITE_FILEPATH not found")
		}
	}

	if "mysql" == dialectName {
		dialectArgs = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True", mysql_user, mysql_passwd, mysql_addr, mysql_database)
	} else if "sqlite3" == dialectName {
		dialectArgs = sqlite_filepath
	}
	db, err := OpenDB()
	if nil != err {
		panic(err)
	}
	defer CloseDB(db)

	db.AutoMigrate(&Application{})
}

func OpenDB() (*gorm.DB, error) {
	return gorm.Open(dialectName, dialectArgs)
}

func CloseDB(_db *gorm.DB) {
	_db.Close()
}
