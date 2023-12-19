package handle

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type nFile struct {
	FileName string `json:"name"` //文件名，充当路径
	FilePath string `json:"path"` //文件路由地址
}

type FileList struct {
	Files []nFile `json:"files"` //文件列表
}

type Server struct {
	GinEngine *gin.Engine
	MList     *FileList
	RList     *FileList
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success"})
}

func CreateServer() *Server {
	GinServer := gin.Default()
	mList := &FileList{}
	rList := &FileList{}
	return &Server{GinServer, mList, rList}
}

func (s *Server) StartServer() {
	s.InitHandler()
	// 启动goroutine运行服务器
	fmt.Println("服务器启动成功，端口号：", viper.GetInt("port"))
	go s.GinEngine.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
}

func CreateNFile(fileName string, filePath string) nFile {
	return nFile{fileName, filePath}
}

func (s *Server) printRoutes(c *gin.Context) {
	for _, route := range s.GinEngine.Routes() {
		fmt.Println(route.Path, route.Method)
	}
	c.JSON(200, gin.H{"message": "success"})
}
