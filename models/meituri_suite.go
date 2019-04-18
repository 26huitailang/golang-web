package models

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Image struct {
	Title string
	Path  string
}
type Suite struct {
	Images []*Image
	Folder string
}

// GetFolderList 获取suite folder的列表
func GetFolderList(path string) []*Suite {
	suiteList := make([]*Suite, 0)
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			suite := &Suite{Folder: file.Name()}
			suiteList = append(suiteList, suite)
		} else {
			continue
		}
	}
	return suiteList
}

// GetFileList 获取路径下jpg的列表
func (suite *Suite) GetFileList(baseFolder string) *Suite {
	imageList := make([]*Image, 0)
	suitePath := filepath.Join(baseFolder, suite.Folder)
	files, _ := ioutil.ReadDir(suitePath)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), "jpg") {
			imagePath := filepath.Join(suite.Folder, file.Name())
			image := &Image{Title: file.Name(), Path: imagePath}
			imageList = append(imageList, image)
		}
	}
	suite.Images = imageList
	return suite
}
