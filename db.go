package gomon

import (
   _ "github.com/jinzhu/gorm/dialects/mysql"
   "github.com/jinzhu/gorm"
   "fmt"
   "log"
)

const (
    user = "root"
    password = ""
    host = "127.0.0.1"
    port = "3306"
)

func MyConn(database string) *gorm.DB {
    connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4", user, password, host, port, database)
    db, err := gorm.Open("mysql", connArgs)
    if err != nil {
        log.Fatal(err)
    }
    return db
}

