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
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var writeLock sync.Mutex

func startServer() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	listener, _ := net.ListenTCP("tcp", addr)
	defer listener.Close()
	fmt.Println("server start and waiting ...")
	go closeListener(listener)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("AcceptTCP err", err)
			break
		}
		msgChan := make(chan *Payload)
		fmt.Println("a new connection:" + conn.RemoteAddr().String())
		go serverListen(conn, msgChan)
		go serverSend(conn, msgChan)
	}
}

func closeListener(tcpListener *net.TCPListener) {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-ch
	tcpListener.Close()
}

func serverListen(conn *net.TCPConn, msgChan chan *Payload) {
	for {
		payload, err := Read(conn)
		if err != nil {
			fmt.Errorf("read from conn err:%s", err)
			return
		}
		fmt.Println(msgChan)
		if payload.P == OriginalClose {
			msgChan <- payload
			conn.Close()
			close(msgChan)
			break
		}
		msgChan <- payload
	}
}

func serverSend(conn *net.TCPConn, msgChan chan *Payload) {
	for v := range msgChan {
		if v == nil {
			break
		}
		Write(&Payload{P: fmt.Sprintf("processed:%s", v.P)}, conn, &writeLock)
	}
}

func main() {
	startServer()
}