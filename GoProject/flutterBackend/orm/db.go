package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "os"
)

var Db *gorm.DB
var err error
func InitDB() {
    dsn := os.Getenv("MSQL_DNS")
    Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    Db.AutoMigrate(&Tbl_User{})
}