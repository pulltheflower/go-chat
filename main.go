package main

import (
	"go-chat/initializer"
	"go-chat/route"
)

func main() {
	initializer.InitConfig()
	initializer.InitMysql()

	r := route.Router()
	r.Run()
}
