package main

import (
	"fmt"
	"html/template"
	"net/http"
	"trans-backend-golang/src/config"
	"trans-backend-golang/src/handle"
	"trans-backend-golang/src/utils"
)

func main() {
	utils.InitDir("trans-data")
	config.InitConfig()
	utils.CreateDataDir()

	server := handle.CreateServer()
	t := template.Must(template.New("files").ParseFiles("templates/files.tmpl"))
	server.GinEngine.SetHTMLTemplate(t)
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", fs)

	go http.ListenAndServe(":3000", nil)
	go server.StartServer()
	fmt.Println("Server is running on port 18080")
	fmt.Println("GUI is running on port 3000")
	<-make(chan error)
}
