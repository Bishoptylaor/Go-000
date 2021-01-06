package demo

import (
	test "github.com/gogo/protobuf/test/example"
	"sync"
	"sync/atomic"
	"testing"
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
 @Time    : 2020/12/18
 @Author  : bishop
 @Software: GoLand
 @Description:
*/

type Config struct {
	a []int
}

func (c *Config) T() {}

func BanchmarkAtomic(b *testing.B) {
	var v atomic.Value
	v.Store(&Config{})

	go func() {
		i := 0
		for {
			i ++
			cfg := &Config{a: []int{i, i + 1, i + 2}}
			// 全拷贝
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for n := 0 ; n < 4 ; n ++ {
		wg.Add(1)
		go func() {
			for n := 0; n<b.N; n++ {
				// load 出的是个interface
				cfg := v.Load().(*Config)
				cfg.T()
			}
		}()
	}
}