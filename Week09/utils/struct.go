package utils

import (
	"fmt"
	"io"
	"net"
	"sync"
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
 @Time    : 2021/1/28
 @Author  : bishop
 @Software: GoLand
 @Description:
*/

type Payload struct {
	P string `json:"p"`
}

const OriginalClose = "close"
const ProcessedClose = "processed:close"

func Read(conn *net.TCPConn) (*Payload, error) {
	ret := &Payload{}

	var buf []byte
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, fmt.Errorf("read err：%s", err.Error())
	}
	ret.P = string(buf)
	return ret, nil
}

// 序列化RequestResponse，并发送
// 序列化后的结构如下：
//   长度  4字节
//   Serial 4字节
//   PayLoad 变长
func Write(r *Payload, conn *net.TCPConn, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	p := []byte(r.P)
	_, _ = conn.Write(p)
	fmt.Println("sending: " + r.P)
}