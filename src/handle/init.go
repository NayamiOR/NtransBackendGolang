package handle

func (s *Server) InitHandler() {
	r := s.GinEngine
	r.GET("/test", Test)
	r.GET("/files", s.ReturnMountFileList)
	r.POST("/receivefiles", ReceiveFiles)
	r.POST("/uploadfiles", s.UploadFiles)
	r.GET("/printroutes", s.printRoutes)
	r.GET("/messages", s.ReturnMessages)
	r.POST("/uploadmessage", UploadMessage)
	r.POST("/receivemessage", ReceiveMessage)
	r.StaticFile("/favicon.ico", "./favicon.ico")
}
