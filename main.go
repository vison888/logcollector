package main

import (
	"github.com/vison888/logcollector/app"
	"github.com/vison888/logcollector/handler/log"
	"github.com/vison888/logcollector/server"
)

func main() {
	app.Init("")

	log.InitLog()

	server.Start()
}
