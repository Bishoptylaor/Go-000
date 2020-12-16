package Week04

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
 @Time    : 2020/12/16
 @Author  : bishop
 @Software: GoLand
 @Description:
*/

import (
	"log"

	"Go-000/Week04/api"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../../configs")
	err := viper.ReadInConfig()
	if err != nil {
		// load config error
		panic(err)
	}
	srv, cleanup, err := api.InitializeServer()
	defer cleanup()
	if err != nil {
		log.Printf("Init Server error:%v\n", err)
		return
	}

	log.Println("Start Server")
	if err = srv.Run(); err != nil {
		log.Printf("Run Server error:%v\n", err)
		return
	}
}