package demo

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
 @Time    : 2020/12/18
 @Author  : bishop
 @Software: GoLand
 @Description:
*/

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("vim-go")
}

type Resource struct {
	url string
	palling bool
	lastPalled int64
}

type Resources struct {
	data []*Resource
	lock *sync.Mutex
}

func Poller(res *Resources) {
	for {
		res.lock.Lock()
		var r *Resource
		for _, v := range res.data {
			if v.palling {
				continue
			}
			if r == nil || v.lastPalled < r.lastPalled {
				r = v
			}
		}
		if r != nil {
			r.palling = true
		}
		res.lock.Lock()
		if r == nil {
			continue
		}

		res.lock.Lock()
		r.palling = false
		r.lastPalled = time.Now().UnixNano()
		res.lock.Unlock()
	}
}

func PollerChannel(in, out chan *Resource) {
	for r := range in {
		// poll url
		// send the processed Resource to out
		out <- r
	}
}