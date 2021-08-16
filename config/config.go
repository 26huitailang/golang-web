package config

import (
	"encoding/json"
	"fmt"
	"github.com/26huitailang/golang_web/constants"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"reflect"
)

var Config *Configuration

type Configuration struct {
	BasePath    string `json:"base_path"`
	IP          string `json:"ip"`
	DeployLevel int    `json:"deploy_level"`
	Port        string `json:"port"`
	DataPath    string `json:"data_path"`
	MediaPath   string `json:"media_path"`
	UIProgress  *UIProgressConf
}

type UIProgressConf struct {
	Show bool `json:"show"`
}

// todo: use cobra to get config

func init() {
	pwd, _ := os.Getwd()
	Config = &Configuration{
		pwd,
		"0.0.0.0",
		constants.Development,
		":8000",
		"/data",
		"/data/media",
		&UIProgressConf{Show: false},
	}
	Config.initConfiguration()
	fmt.Println("CONFIG:", Config)
}

// 加载自定义配置，覆盖默认配置
func (conf *Configuration) initConfiguration() {
	// 文件是否存在
	file, err := os.Open("config_custom.json")
	if err != nil {
		log.Warn("config_custom.json not existed!\nUse default\n")
		return
	}
	defer file.Close()
	// 读取json为map
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// var customConfig Configuration
	var customMap map[string]interface{}
	json.Unmarshal(data, &customMap)

	// 遍历config结构体，判断是否有覆盖内容
	t := reflect.TypeOf(conf).Elem()
	v := reflect.ValueOf(conf).Elem()
	for i := 0; i < t.NumField(); i++ {
		// 比较tag
		fieldInfo := t.Field(i)
		tag := fieldInfo.Tag.Get("json")
		if value, ok := customMap[tag]; ok {
			log.Printf("tag: [%s %v] replaced by [%v]", tag, v.Field(i).Interface(), value)
			valueSet := reflect.ValueOf(value)
			v.FieldByName(fieldInfo.Name).Set(valueSet.Convert(fieldInfo.Type))
		}
	}
	log.Println("custom config:", conf)
}
