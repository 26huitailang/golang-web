package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/26huitailang/golang_web/constants"
	log "github.com/sirupsen/logrus"
)

var Config *Configuration

type Configuration struct {
	BasePath           string `json:"base_path"`
	ConfigPath         string `json:"config_path"`
	IP                 string `json:"ip"`
	DeployLevel        int    `json:"deploy_level"`
	Port               string `json:"port"`
	DB                 string `json:"db"`
	DataPath           string `json:"data_path"`
	MediaPath          string `json:"media_path"`
	SessionExpiredTime int
	UIProgress         *UIProgressConf
}

type UIProgressConf struct {
	Show bool `json:"show"`
}

// todo: use cobra to get config

func init() {
	pwd := GetCurrentPath()
	projectPath := filepath.Dir(pwd)
	Config = &Configuration{
		BasePath:           projectPath,
		ConfigPath:         pwd,
		IP:                 "0.0.0.0",
		DeployLevel:        constants.Development,
		Port:               ":8000",
		DB:                 "test.db",
		DataPath:           "/data",
		MediaPath:          "/data/media",
		SessionExpiredTime: 3600,
		UIProgress:         &UIProgressConf{Show: false},
	}
	Config.initConfiguration()
	log.Debug("CONFIG:", Config)
}

// 加载自定义配置，覆盖默认配置
func (conf *Configuration) initConfiguration() {
	// 文件是否存在
	customConfPath := filepath.Join(GetCurrentPath(), "config_custom.json")
	file, err := os.Open(customConfPath)
	if err != nil {
		log.Warn(customConfPath, "not existed!\nUse default\n")
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
			log.Debugf("tag: [%s %v] replaced by [%v]", tag, v.Field(i).Interface(), value)
			valueSet := reflect.ValueOf(value)
			v.FieldByName(fieldInfo.Name).Set(valueSet.Convert(fieldInfo.Type))
		}
	}
	log.Debug("custom config:", conf)
}

func GetCurrentPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
