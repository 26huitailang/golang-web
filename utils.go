package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
)

// WalkBasePath 根据配置的路径返回文件结构
// 现在为两层，theme/suites
type themeInfo struct {
	Name  string
	Error error
}

// InitTheme to init database from local files
func InitTheme(conf *Configuration) {
	dir, _ := ioutil.ReadDir(conf.BasePath)
	finish := make(chan themeInfo)
	var wg sync.WaitGroup
	tx := DB.Begin()
	for _, folder := range dir {
		if !folder.IsDir() {
			continue
		}
		wg.Add(1)
		// handle one theme
		go func(folder os.FileInfo) {
			defer wg.Done()
			theme := &Theme{Name: folder.Name()}
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

func initSuiteByTheme(tx *gorm.DB, themePath string, theme *Theme, out chan<- themeInfo) {
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
		suite := &Suite{Name: info.Name(), ThemeID: theme.ID}
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

func initImageBySuite(tx *gorm.DB, suite *Suite, suitePath string) (n int, err error) {
	n = 0
	err = filepath.Walk(suitePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".jpg") {
			themeName := filepath.Base(filepath.Dir(suitePath))
			imgPath := filepath.Join(themeName, filepath.Base(suitePath), info.Name())
			img := &Image{Path: imgPath, SuiteID: suite.ID}
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

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
