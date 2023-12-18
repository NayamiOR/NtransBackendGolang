package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// 初始化目录，如果目录不存在，则创建目录
func InitDir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.Mkdir(dirName, os.ModePerm)
	}
}

// 初始化文件，如果文件不存在，则创建文件
func InitFile(fileName string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		os.Create(fileName)
	}
}

// 从字节流保存文件到本地
func SaveFile(file []byte, fileName string) error {
	rootPath := viper.GetString("download_location")
	return SaveFileWithDir(file, fileName, rootPath)
}
func SaveFileWithDir(file []byte, fileName string, dirName string) error {
	rootPath := "trans-data/" + dirName
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		os.Mkdir(rootPath, os.ModePerm)
	}
	filePath := rootPath + fileName
	f, err := os.Create(filePath)
	if err != nil {
		log.Println("文件", fileName, "创建失败，错误信息：", err)
		return err
	}
	defer f.Close()
	_, err = f.Write(file)
	if err != nil {
		log.Println("文件", fileName, "写入失败")
		log.Println(err)
		return err
	}
	return nil
}

// 创建数据目录
func CreateDataDir() {
	InitDir(fmt.Sprintf("trans-data/%s", viper.GetString("download_location")))
	InitDir(fmt.Sprintf("trans-data/%s", viper.GetString("mount_location")))
	InitFile("trans-data/messages.txt")
	InitFile("trans-data/mounted-files/messages.txt")
}
