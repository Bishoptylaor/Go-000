package data_race

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
 @Time    : 2020/12/9
 @Author  : bishop
 @Software: GoLand
 @Description:
*/

import "fmt"

type IceCreamMaker interface {
	Hello()
}
type Ben struct {
	id int
	name string
}
func (b *Ben) Hello() {
	fmt.Printf("Ben says, hello my name is %s \n", b.name)
}
type John struct {
	name string
}
func (j *John) Hello() {
	fmt.Printf("John says, hello my name is %s \n", j.name)
}

func main() {
	var ben = &Ben{id: 1, name: "Ben"}
	var john = &John{name: "John"}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()
	loop0 = func() {
		maker = ben
		go loop1()
	}
	loop1 = func() {
		maker = john
		go loop0()
	}

	go loop0()
	for {
		maker.Hello()
	}
}