package main

import (
	"github.com/ushanovsn/tgbotmon/internal/server"
	"time"
)

func main() {
	srv := server.InitServer()
	server.StartServer(srv)

	time.Sleep(3 * time.Second)

	server.StopServer(srv)
}
