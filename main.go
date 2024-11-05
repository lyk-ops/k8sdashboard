package main

import (
	"kubeimook/api/initallize"
	"kubeimook/global"
)

func main() {
	r := initallize.Router()
	initallize.Viper()
	initallize.K8s() //初始化客户端
	panic(r.Run(global.CONF.System.Addr))
}
