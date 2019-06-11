package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/26huitailang/golang-web/database"
	"github.com/26huitailang/golang-web/models"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Config *Configuration
var db = database.DB

type Configuration struct {
	BasePath    string `json:"base_path"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
	DeployLevel int    `json:"deploy_level"`
}

func init() {
	Config = &Configuration{}
	Config.initConfiguration()
	fmt.Printf("COFIG: %v", Config)
}

// 加载默认配置config.json
func (conf *Configuration) initConfiguration() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Error opening JSON file:", err)
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error reading JSON data:", err)
	}
	json.Unmarshal(jsonData, &conf)
	conf.initCustomConfig()
	log.Println("config:", conf)
}

// 加载自定义配置，覆盖默认配置
func (conf *Configuration) initCustomConfig() {
	// 文件是否存在
	file, err := os.Open("config_custom.json")
	if err != nil {
		log.Warn("config_custom.json not existed!\nUse default config.json\n")
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
			v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
		}
	}
}

type themeInfo struct {
	Name  string
	Error error
}

// InitTheme 根据配置的路径返回文件结构
// 现在为两层，theme/suites
// InitTheme to init database from local files
func (conf *Configuration) InitTheme() {
	dir, _ := ioutil.ReadDir(conf.BasePath)
	finish := make(chan themeInfo)
	var wg sync.WaitGroup
	tx := db.Begin()
	for _, folder := range dir {
		if !folder.IsDir() {
			continue
		}
		wg.Add(1)
		// handle one theme
		go func(folder os.FileInfo) {
			defer wg.Done()
			theme := &models.Theme{Name: folder.Name()}
			tx.Create(theme)
			themePath := filepath.Join(conf.BasePath, folder.Name())
			initSuiteByTheme(tx, themePath, theme, finish)
			// Theme[folder.Name()] = suites
		}(folder)
	}
	go func() {
		wg.Wait()
		close(finish)
	}()

	for tInfo := range finish {
		if tInfo.Error != nil {
			tx.Rollback()
			panic(tInfo.Error)
		}
		log.Println("theme inited:", tInfo.Name)
	}
	tx.Commit()
}

func initSuiteByTheme(tx *gorm.DB, themePath string, theme *models.Theme, out chan<- themeInfo) {
	err := filepath.Walk(themePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			return nil
		}
		if filepath.Base(themePath) == info.Name() {
			return nil
		}
		suite := &models.Suite{Name: info.Name(), ThemeID: theme.ID}
		tx.Create(suite)
		suitePath := filepath.Join(themePath, info.Name())
		n, _ := initImageBySuite(tx, suite, suitePath)
		log.Printf("Suite: %s | Image: %d insert", suite.Name, n)
		return nil
	})
	out <- themeInfo{
		theme.Name,
		err,
	}
}

func initImageBySuite(tx *gorm.DB, suite *models.Suite, suitePath string) (n int, err error) {
	n = 0
	err = filepath.Walk(suitePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".jpg") {
			themeName := filepath.Base(filepath.Dir(suitePath))
			imgPath := filepath.Join(themeName, filepath.Base(suitePath), info.Name())
			img := &models.Image{Path: imgPath, SuiteID: suite.ID}
			tx.Create(img)
			n++
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return n, err
	}
	return n, nil
}
