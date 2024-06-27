package main

import (
	"time"
	"github.com/ushanovsn/tgbotmon/internal/server"
)



func main() {
	srv := server.InitServer()
	server.StartServer(srv)

	time.Sleep(3 * time.Second)

	server.StopServer(srv)
}