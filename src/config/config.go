package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./trans-data")
	// check if config file exists
	// if not, create one
	if _, err := os.Stat("trans-data/config.json"); os.IsNotExist(err) {
		file, _ := os.Create("trans-data/config.json")
		file.Close()
		viper.SetDefault("download_location", "downloaded-files/")
		viper.SetDefault("port", "18080")
		viper.SetDefault("mount_location", "mounted-files/")
		viper.WriteConfig()
	} else {
		viper.ReadInConfig()
	}
}
func GetPropsList() map[string]string {
	properties := make(map[string]string)
	return properties
}

func GetProp(key string) string {
	properties := GetPropsList()
	return properties[key]
}

func UpdateProp(key string, value string) {
	properties := GetPropsList()
	properties[key] = value

}
