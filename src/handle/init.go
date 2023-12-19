package handle

func (s *Server) InitHandler() {
	r := s.GinEngine
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.GET("/test", Test)
	r.GET("/printroutes", s.printRoutes)
	r.GET("/files", s.ShowMountedFiles)
	r.GET("/filesReceived", s.ShowReceivedFiles)
	r.GET("/messages", ReturnMessages)
	r.GET("/messagesreceived", ReturnMessagesReceived)
	r.POST("/uploadfiles", s.UploadFiles)
	r.POST("/receivefiles", s.ReceiveFiles)
	r.POST("/uploadmessage", UploadMessage)
	r.POST("/receivemessage", ReceiveMessage)
}
