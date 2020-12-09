package Never_start_a_goroutine_without_knowning_when_it_will_stop

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
 @Time    : 2020/12/8
 @Author  : bishop
 @Software: GoLand
 @Description:
*/

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// v1
func main1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "Hello")
	})

	go http.ListenAndServe("127.0.0.1:8081", http.DefaultServeMux)
	http.ListenAndServe("0.0.0.0:8080", mux)
}

// v2
func serverApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "Hello")
	})

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}
func serverDebug() {
	if err := http.ListenAndServe("127.0.0.1:8081", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}
func main2() {
	go serverDebug()
	go serverApp()
	select {

	}
}