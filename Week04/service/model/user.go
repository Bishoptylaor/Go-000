package model

import (
	"Go-000/Week04/service/model/dao"
	"context"
	"database/sql"
)

type UserDB interface {
}

type user struct {
	db *sql.DB
}

/*
------------pub结构------------
------------数据库结构(desc)------------
*/

type User struct {
	Ct int64 `json:"ct"`
	Ut int64 `json:"ut"`
	Uid int64 `json:"uid"`
	Name string `json:"name"`
	Headimg string `json:"headimg"`
}

func (user) DecodeFromDB(ctx context.Context, logs []*dao.TestcaseUser) []*User {
	var items []*User
	for _, l := range logs {
		items = append(items, &User{
			Ct:      l.Ct,
			Ut:      l.Ut,
			Uid:     int64(l.UID),
			Name:    l.Name,
			Headimg: l.Headimg,
		})
	}
	return items
}

func (u *user) Insert(ctx context.Context, data map[string]interface{}) (int64, error) {
	return dao.InsertTestcaseUser(ctx, u.db, []map[string]interface{}{data})
}

func (u *user) QueryList(ctx context.Context, ct int64) ([]*dao.TestcaseUser, error) {
	where := map[string]interface{}{
		"ct >":        ct,
	}
	return dao.GetMultiTestcaseUser(ctx, u.db, where)
}

func (u *user) QueryOne(ctx context.Context, uid int64) (*dao.TestcaseUser, error) {
	where := map[string]interface{}{
		"uid":        uid,
	}
	return dao.GetOneTestcaseUser(ctx, u.db, where)
}
