package dao

/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2020/12/16
 @Author  : bishop
 @Software: GoLand
 @Description: dao层实现
*/

import (
	"context"
	"database/sql"
)

const TestcaseUserTable = "testcase_user"

// TestcaseUser is a mapping object for testcase_user table in mysql
type TestcaseUser struct {
	UID     uint64 `json:"uid" bdb:"uid"`         // 自增id
	Ct      int64  `json:"ct" bdb:"ct"`           // 创建时间
	Ut      int64  `json:"ut" bdb:"ut"`           // 更新时间
	Name    string `json:"name" bdb:"name"`       // 用户名
	Headimg string `json:"headimg" bdb:"headimg"` // 头像
}

//GetOneTestcaseUser gets one record from table testcase_user by condition "where"
func GetOneTestcaseUser(ctx context.Context, db *sql.DB, where map[string]interface{}) (res *TestcaseUser, err error) {
	return
}

//GetMultiTestcaseUser gets multiple records from table testcase_user by condition "where"
func GetMultiTestcaseUser(ctx context.Context, db *sql.DB, where map[string]interface{}) (res []*TestcaseUser, err error) {
	return
}

//InsertTestcaseUser inserts an array of data into table testcase_user
func InsertTestcaseUser(ctx context.Context, db *sql.DB, data []map[string]interface{}) (rid int64, err error) {
	return
}

//UpdateTestcaseUser updates the table testcase_user
func UpdateTestcaseUser(ctx context.Context, db *sql.DB, where, data map[string]interface{}) (rid int64, err error) {
	return
}

// DeleteTestcaseUser deletes matched records in testcase_user
func DeleteTestcaseUser(ctx context.Context, db *sql.DB, where map[string]interface{}) (rid int64, err error) {
	return
}

// GetCountOfTestcaseUser get number from testcase_user by condition "where"
func GetCountOfTestcaseUser(ctx context.Context, db *sql.DB, where map[string]interface{}) (count int, err error) {
	return
}