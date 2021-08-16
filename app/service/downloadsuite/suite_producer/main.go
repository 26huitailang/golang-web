package main

import (
	"flag"
	"github.com/26huitailang/golang_web/app/service/downloadsuite"
	"github.com/nsqio/go-nsq"
	"log"
)

var URL string
var folder string

func init() {
	// https://www.lanvshen.com/x/86/
	flag.StringVar(&URL, "url", "", "Theme首页")
	flag.StringVar(&folder, "folder", "tmp", "保存的suite路径")
}

func produce(url, folderSave string) {
	theme := downloadsuite.NewTheme(url, folderSave)
	cfg := nsq.NewConfig()
	nsqAddr := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(nsqAddr, cfg)
	if err != nil {
		log.Fatal(err)
	}

	topic := "mtr_suite"
	if err = theme.Produce(producer, topic); err != nil {
		log.Fatalf("produce error: %s", err.Error())
	}
}

func main() {
	flag.Parse()
	produce(URL, folder)
}
