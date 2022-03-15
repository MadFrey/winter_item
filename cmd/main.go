package main

import (
	"blog/dao"
	"blog/router"
	"blog/util"
	"log"
)

func main() {
	dns := "root:sjk123456@tcp(localhost:3306)/blog?parseTime=true"
	err := dao.Init(dns)
	if err != nil {
		log.Println(err)
		return
	}
	err = util.InitClient()
	if err != nil {
		log.Println(err)
		return
	}
	r := router.CreateRouter()
	r.Run(":9090")
}
