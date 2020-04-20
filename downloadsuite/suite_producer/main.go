package main

import (
	"github.com/nsqio/go-nsq"
	"golang_web/downloadsuite"
	"log"
)

func produce() {
	theme := downloadsuite.NewTheme("https://www.lanvshen.com/x/98/", "~/Downloads/mtr")
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
	produce()
}
