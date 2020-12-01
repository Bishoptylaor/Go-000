package Week02

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

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
 @Time    : 2020/11/28
 @Author  : bishop
 @Software: GoLand
 @Description:
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

*/
func main() {
	_, err := TestQuery()
	fmt.Println(err)
}

type User struct {
	Uid int64 `json:"uid"`
	HeadImg string `json:"head_img"`
}

type dao struct {}
func (dao) QueryUser(id int) (u *User, err error) {
	if id == 0 {
		return nil, sql.ErrNoRows
	}
	return &User{Uid: 10032, HeadImg: "https://avatars3.githubusercontent.com/u/19482929?s=60&v=4"}, nil
}

func TestQuery() (img string, err error) {
	d := dao{}
	user, err := d.QueryUser(123)
	if err != nil && errors.Is(err, sql.ErrNoRows){
		return "", errors.New("Can`t Find This User")
	} else if err != nil {
		return "", errors.Wrap(err,"Internal Error")
	}
	return user.HeadImg, nil
}