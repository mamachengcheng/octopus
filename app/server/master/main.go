package main

import (
	"log"
	"master/manager/api"
	"master/models"
	"time"
)

func main() {
	//gin.SetMode(gin.DebugMode)
	//
	//if cfg, err := ini.Load(resource.ConfFilePath); err != nil {
	//	panic(err)
	//} else {
	//	server := cfg.Section("server")
	//	address := server.Key("http").String() + ":" + server.Key("port").String()
	//
	//	router := gin.Default()
	//	routers.InitRouter(router)
	//	router.Run(address)
	//}
	for {
		for _, ip := range models.GetIPList() {


			nodeInfo := api.GetNodeInfo(ip)
			models.SetNodeInfo(ip, nodeInfo)
			time.Sleep(time.Second)
			log.Printf("end %v", ip)
		}
	}
}