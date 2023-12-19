package handle

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"
	"trans-backend-golang/src/utils"

	"github.com/gin-gonic/gin"
)

func ReturnMessages(c *gin.Context) {
	utils.InitFile("trans-data/mounted-files/messages.txt")
	// read messages.txt
	// content, _ := os.ReadFile("trans-data/mounted-files/messages.txt")
	// c.JSON(200, gin.H{"messages": string(content)})
	c.File("trans-data/mounted-files/messages.txt")
}

func (s *Server) ReturnMountFileList(c *gin.Context) {
	json, err := json.Marshal(s.MList.Files)
	if err != nil {
		log.Println("json转换失败，错误信息：", err)
	}
	c.JSON(200, gin.H{"files": string(json)})
}

func (s *Server) ShowMountedFiles(c *gin.Context) {
	fileNames := make([]string, 0)
	for _, file := range s.MList.Files {
		fileNames = append(fileNames, file.FileName)
	}
	c.HTML(200, "files.tmpl", fileNames)
}

/*
- 接收文件的路由
- 注意要使用encode="multipart/form-data"
- <form action="http://localhost:8080/uploadfile" method="post" enctype="multipart/form-data">
  - <input type="file" id="file" name="file" multiple><br><br>
  - <input type="submit" value="Submit">

- </form>
*/
func ReceiveFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("文件获取失败,错误信息：", err)
		return
	}
	files := form.File["file"]
	for _, file := range files {
		fileContent, err := file.Open()
		// 用utf-8编码读取文件名
		fileName := html.UnescapeString(file.Filename)

		if err != nil {
			log.Println("文件", file.Filename, "打开失败,错误信息：", err)
			return
		}
		bytes := make([]byte, file.Size)
		fileContent.Read(bytes)

		utils.SaveFile(bytes, fileName)
	}

	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	log.Println(len(files), "个文件上传成功")
}

/**
 * 接收消息的路由
 * <input type="input" name="text"><br><br>
 */
func ReceiveMessage(c *gin.Context) {
	request := c.Request
	text := request.FormValue("text")
	// 创建messages.txt文件，如果文件已存在，则打开该文件，如果文件不存在，则创建该文件
	messagesFile, err := os.OpenFile("trans-data/messages.txt", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("文件打开失败，错误信息：", err)
		return
	}
	defer messagesFile.Close()

	messagesFile.WriteString("--------------------\n")
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	messagesFile.WriteString(timestamp + ":\n")
	messagesFile.WriteString(text + "\n")

	log.Println("接收到消息：", text)
	c.JSON(200, gin.H{"message": "success"})
}

// 发送文件
func (s *Server) MountFile(file nFile) {
	for _, f := range s.MList.Files {
		if f.FilePath == file.FilePath {
			log.Println("文件", file.FileName, "已存在")
			return
		}
	}
	s.MList.Files = append(s.MList.Files, file)
	s.GinEngine.GET(fmt.Sprintf("/files/%s", file.FilePath), func(c *gin.Context) {
		c.File(fmt.Sprintf("trans-data/mounted-files/%s", file.FileName))
	})
}

func (s *Server) UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("文件获取失败,错误信息：", err)
		return
	}
	files := form.File["file"]
	nFs := make([]nFile, 0)

	for _, file := range files {
		fileContent, err := file.Open()
		// 用utf-8编码读取文件名
		fileName := html.UnescapeString(file.Filename)

		if err != nil {
			log.Println("文件", file.Filename, "打开失败,错误信息：", err)
			return
		}
		bytes := make([]byte, file.Size)
		fileContent.Read(bytes)
		utils.SaveFileWithDir(bytes, fileName, "mounted-files/")
		nF := CreateNFile(fileName, fileName)
		nFs = append(nFs, nF)
	}

	for _, file := range nFs {
		s.MountFile(file)
	}
	c.JSON(200, gin.H{"message": fmt.Sprintf("%d files uploaded!", len(files))})
}

func UploadMessage(c *gin.Context) {
	message := c.PostForm("message")
	messageFile, err := os.OpenFile("trans-data/mounted-files/messages.txt", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("文件打开失败，错误信息：", err)
		return
	}
	defer messageFile.Close()
	messageFile.WriteString("--------------------\n")
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	messageFile.WriteString(timestamp + ":\n")
	messageFile.WriteString(message + "\n")
	c.JSON(200, gin.H{"message": "success"})
}
func ReturnMessagesReceived(c *gin.Context) {
	utils.InitFile("trans-data/messages.txt")
	c.File("trans-data/messages.txt")
}
