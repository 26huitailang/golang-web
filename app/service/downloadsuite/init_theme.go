package downloadsuite

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/26huitailang/golang_web/app/model"
	"github.com/26huitailang/golang_web/config"
	"github.com/26huitailang/golang_web/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type themeInfo struct {
	Name  string
	Error error
}

// InitTheme 根据配置的路径返回文件结构
// 现在为两层，theme/suites
// InitTheme to init database from local files
// todo: 跳过已经存在的suite
func InitTheme(conf *config.Configuration) {
	dir, _ := ioutil.ReadDir(conf.MediaPath)
	finish := make(chan themeInfo)
	var wg sync.WaitGroup
	db := database.NewDatabaseStore().DB()
	tx := db.Begin()
	for _, folder := range dir {
		if !folder.IsDir() {
			continue
		}
		wg.Add(1)
		// handle one theme
		go func(folder os.FileInfo) {
			defer wg.Done()
			theme := &model.Theme{Name: folder.Name()}
			tx.Create(theme)
			themePath := filepath.Join(conf.MediaPath, folder.Name())
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
		logrus.Println("theme inited:", tInfo.Name)
	}
	tx.Commit()
}

func initSuiteByTheme(tx *gorm.DB, themePath string, theme *model.Theme, out chan<- themeInfo) {
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
		suite := &model.Suite{Name: info.Name(), ThemeID: theme.ID}
		// 跳过已存在的suite
		if err := tx.Create(suite).Error; err != nil {
			logrus.Infoln("db instert downloadsuite skip existed:", suite.Name)
			return nil
		}
		suitePath := filepath.Join(themePath, info.Name())
		n, _ := initImageBySuite(tx, suite, suitePath)
		logrus.Printf("Suite: %s | Image: %d insert", suite.Name, n)
		return nil
	})
	out <- themeInfo{
		theme.Name,
		err,
	}
}

func initImageBySuite(tx *gorm.DB, suite *model.Suite, suitePath string) (n int, err error) {
	n = 0
	err = filepath.Walk(suitePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".jpg") {
			themeName := filepath.Base(filepath.Dir(suitePath))
			imgPath := filepath.Join(themeName, filepath.Base(suitePath), info.Name())
			img := &model.Image{Path: imgPath, SuiteID: suite.ID}
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
