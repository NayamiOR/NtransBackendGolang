package main

import (
	"fmt"
	"trans-backend-golang/src/config"
	"trans-backend-golang/src/handle"
	"trans-backend-golang/src/utils"
)

func main() {
	utils.InitDir("trans-data")
	config.InitConfig()
	utils.CreateDataDir()
	server := handle.CreateServer()
	fmt.Println("Server is running on port 18080")
	fmt.Println("GUI is running on port 3000")
	// c := cli.CreateCLI(server.GinEngine)
	// go c.StartCLI()

	go server.StartServer()
	<-make(chan error)
}
