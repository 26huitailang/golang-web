package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"golang_web/constants"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"golang_web/database"
	"golang_web/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Config *Configuration
var db = database.DB

type Configuration struct {
	BasePath    string `json:"base_path"`
	IP          string `json:"ip"`
	DeployLevel int    `json:"deploy_level"`
	Port        string `json:"port"`
}

func init() {
	flag.Parse()
	pwd, _ := os.Getwd()
	Config = &Configuration{
		pwd,
		"0.0.0.0",
		constants.Development,
		":8000",
	}
	Config.initConfiguration()
	fmt.Println("CONFIG:", Config)
}

// 加载自定义配置，覆盖默认配置
func (conf *Configuration) initConfiguration() {
	// 文件是否存在
	file, err := os.Open("config_demo.json")
	if err != nil {
		log.Warn("config_custom.json not existed!\nUse default config_demo.json\n")
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

type themeInfo struct {
	Name  string
	Error error
}

// InitTheme 根据配置的路径返回文件结构
// 现在为两层，theme/suites
// InitTheme to init database from local files
// todo: 跳过已经存在的suite
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
		// 跳过已存在的suite
		if err := tx.Create(suite).Error; err != nil {
			log.Infoln("db instert downloadsuite skip existed:", suite.Name)
			return nil
		}
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
