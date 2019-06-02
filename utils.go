package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// WalkBasePath 根据配置的路径返回文件结构
// 现在为两层，theme/suites
type themeInfo struct {
	Name   string
	Suites []*Suite
}

func InitTheme(conf *Configuration) {
	dir, _ := ioutil.ReadDir(conf.BasePath)
	themeCh := make(chan themeInfo, 10)
	var wg sync.WaitGroup
	for _, folder := range dir {
		if !folder.IsDir() {
			continue
		}
		wg.Add(1)
		go func(folder os.FileInfo) {
			defer wg.Done()
			themePath := filepath.Join(conf.BasePath, folder.Name())
			GetSuiteByTheme(themePath, folder.Name(), themeCh)
			// Theme[folder.Name()] = suites
		}(folder)
	}
	go func() {
		wg.Wait()
		close(themeCh)
	}()

	for item := range themeCh {
		Theme[item.Name] = item.Suites
	}
	fmt.Println(Theme)
}

// 遍历文件夹，返回该目录下所有的文件夹，不包含再下层
func GetSuiteByTheme(themePath string, folderName string, out chan<- themeInfo) {
	ret := []*Suite{}
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
		images, _ := GetImageFiles(filepath.Join(themePath, info.Name()))
		suite := &Suite{Name: info.Name(), Images: images}
		ret = append(ret, suite)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	out <- themeInfo{
		folderName,
		ret,
	}
}

func GetImageFiles(suitePath string) (ret []Image, err error) {
	err = filepath.Walk(suitePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".jpg") {
			imgPath := filepath.Join(filepath.Base(filepath.Dir(suitePath)), filepath.Base(suitePath), info.Name())
			img := Image{imgPath}
			ret = append(ret, img)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ret, nil
}
