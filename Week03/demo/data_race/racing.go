package data_race

import (
	"fmt"
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
 @Time    : 2020/12/8
 @Author  : bishop
 @Software: GoLand
 @Description:
*/

var Wait sync.WaitGroup
var Counter int = 0

func main() {
	for routine := 1 ; routine <= 2; routine ++ {
		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Println("Final Counter:", Counter)
}

func Routine(id int) {
	for count := 0 ; count < 2; count ++ {
		value := Counter
		value ++
		Counter = value
	}
	Wait.Done()
}