package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"golang_web/downloadsuite"
	"log"
)

type myMessageHandler struct{}

func processMessage(body []byte) (*downloadsuite.SuiteInfo, error) {
	log.Printf("%s", body)
	mtr := new(downloadsuite.SuiteInfo)
	err := json.Unmarshal(body, mtr)
	return mtr, err
}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}
	suiteInfo, err := processMessage(m.Body)
	if err != nil {
		log.Fatalf("unmarshal error: %s", err.Error())
	}

	// todo: 消息应该收到mtr suite的初始化消息
	mtr := downloadsuite.NewMeituriSuite(suiteInfo.FirstPage, suiteInfo.FolderPath, downloadsuite.MeituriParser{})
	cfg := nsq.NewConfig()
	nsqAddr := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(nsqAddr, cfg)
	if err != nil {
		log.Printf("produce error: %s", err.Error())
		return err
	}

	topic := "mtr_image"
	if err = mtr.Produce(producer, topic); err != nil {
		log.Printf("produce error: %s", err.Error())
		return err
	}

	return nil
}

func consume() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("mtr_suite", "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(&myMessageHandler{})

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	// err = consumer.ConnectToNSQD("localhost:4150")
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the consumer.
	// consumer.Stop()
	<-consumer.StopChan
}

func main() {
	consume()
}
