package main

import (
	"blog/dao"
	"blog/router"
	"log"
)


func main() {
	dns:="root:sjk123456@tcp(localhost:3306)/blog?parseTime=true"
	err:=dao.Init(dns)
	if err != nil {
		log.Fatal(err)
	}
	r := router.CreateRouter()
	r.Run(":9090")
}

