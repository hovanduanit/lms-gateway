package main

import "github.com/hovanduanit/lms-proto/servers"

func main() {
	ginServer := servers.NewServer()
	ginServer.Init()
	ginServer.Serve()
}
