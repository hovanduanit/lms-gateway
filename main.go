package main

import "github.com/hovanduanit/lms-gateway/servers"

func main() {
	ginServer := servers.NewServer()
	ginServer.Init()
	ginServer.Serve()
}
