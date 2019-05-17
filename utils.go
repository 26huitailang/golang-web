package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// WalkBasePath 根据配置的路径返回文件结构
// 现在为两层，theme/suites
func InitTheme(conf *Configuration) {
	dir, _ := ioutil.ReadDir(conf.BasePath)
	for _, folder := range dir {
		if !folder.IsDir() {
			continue
		}
		themePath := filepath.Join(conf.BasePath, folder.Name())
		suites, _ := GetSuiteByTheme(themePath)
		Theme[folder.Name()] = suites
	}
	fmt.Println(Theme)
}

// 遍历文件夹，返回该目录下所有的文件夹，不包含再下层
func GetSuiteByTheme(themePath string) (ret []*Suite, err error) {
	err = filepath.Walk(themePath, func(path string, info os.FileInfo, err error) error {
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
		return nil, err
	}
	return ret, nil
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
