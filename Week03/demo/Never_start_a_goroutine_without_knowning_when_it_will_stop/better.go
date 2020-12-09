package Never_start_a_goroutine_without_knowning_when_it_will_stop

import (
	"context"
	"fmt"
	"net/http"
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
 @Time    : 2020/12/8
 @Author  : bishop
 @Software: GoLand
 @Description:
*/

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{Addr: addr, Handler: handler}

	go func() {
		<-stop
		_ = s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}
func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		// done <- serverDebug(stop)
	}()
	go func() {
		// done <- serverApp(stop)
	}()

	var stopped bool
	for i := 0 ;i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error:", err)
		}
		if !stopped {
			stopped = false
			close(stop)
		}
	}
}