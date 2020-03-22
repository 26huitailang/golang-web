/*
Package downloadsuite to download one downloadsuite images.

	Now support meituri.
*/
package downloadsuite

import (
	"errors"
	"strconv"
)

// ISuiteOperator interface 定义了Suite的操作集合
type ISuiteOperator interface {
	//GetPageURLs(chan string)
	//GetImgURLs(chPage <-chan string, chFailedImg <-chan string) <-chan string
	Download(isTheme bool)
	//GetOrgURL() string
}

type Suite struct {
	Operator ISuiteOperator
}

func NewSuite(operator ISuiteOperator) *Suite {
	return &Suite{
		Operator: operator,
	}
}

// todo: can this work
// NewSuiteV2 to be a factory to produce suite manager depends on url host
func NewSuiteV2(url string) (ISuite, error) {
	switch url {
	case "mtr":
		return NewSuiteDemo(url), nil
	default:
		return nil, errors.New("now")
	}
}

type ISuite interface {
	Parse() // parse the image urls/title...
	Download()
}

type SuiteDemo struct {
	images []string
	title  string
	url    string
}

func NewSuiteDemo(url string) ISuite {
	return &SuiteDemo{url: url}
}
func (sd *SuiteDemo) Parse() {
	println("parse")
	for i := 0; i < 10; i++ {
		sd.images = append(sd.images, strconv.Itoa(i))
	}
}

func (sd *SuiteDemo) Download() {
	println("download")
	for _, item := range sd.images {
		println(item)
	}
}
