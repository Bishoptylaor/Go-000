package model

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const (
	cluster = "TESTCASE"
)

const (
	StatusScript  = 1  // 草稿
	StatusOffline = 2  // 未开始
	StatusOnline  = 3  // 进行中
	StatusEnd     = 4  // 已结束
	StatusDel     = 10 // 已删除
)

func InitModel() error {
	// connect Db

	return nil
}

var Provider = wire.NewSet(NewDB, NewUser)

func NewUser(db *sql.DB) UserDB {
	return &user{db: db}
}


func NewDB() (db *sql.DB, cleanup func(), err error) {
	db, err = sql.Open("mysql", viper.GetString("mysql.dsn"))
	cleanup = func() {
		if err == nil {
			db.Close()
		}
	}
	return
}

