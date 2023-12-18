package main

import (
	"trans-backend-golang/src/config"
	"trans-backend-golang/src/handle"
	"trans-backend-golang/src/utils"
)

func main() {
	utils.InitDir("trans-data")
	config.InitConfig()
	utils.CreateDataDir()
	server := handle.CreateServer()
	// c := cli.CreateCLI(server.GinEngine)
	// go c.StartCLI()

	go server.StartServer()
	<-make(chan error)
}
