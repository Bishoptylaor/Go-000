package main

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
 @Time    : 2021/1/27
 @Author  : bishop
 @Software: GoLand
 @Description:
*/
import (
	. "Go-000/Week09/utils"
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net"
	"sync"
)

var writeLock sync.Mutex

func startClient(msgChan chan string) error {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return err
	}

	go clientListen(conn)
	go clientSend(conn, msgChan)
	return nil
}

func clientListen(conn *net.TCPConn) {
	for {
		r, err := Read(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if r.P == ProcessedClose {
			conn.Close()
			break
		}
		fmt.Println("receiving:" + r.P)
	}
}

func clientSend(conn *net.TCPConn, msgChan chan string) {
	for v := range msgChan {
		Write(&Payload{P: v}, conn, &writeLock)
	}
}

func main() {
	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)
	msgChan := make(chan string, 10)
	group.Go(func() error {
		return startClient(msgChan)
	})
	group.Go(func() error {
		for {
			var payload string
			fmt.Scanln(&payload)
			msgChan <- payload
			if payload == OriginalClose {
				fmt.Println("received close command")
				break
			}
		}
		return errors.New("exit")
	})
	err := group.Wait()
	fmt.Println(err)
}