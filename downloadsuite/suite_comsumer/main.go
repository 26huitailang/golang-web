package main

import (
	"github.com/nsqio/go-nsq"
	"golang_web/downloadsuite"
	"log"
)

type myMessageHandler struct{}

func processMessage(body []byte) error {
	log.Printf("%s", body)
	return nil
}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}
	err := processMessage(m.Body)

	// todo: 消息应该收到mtr suite的初始化消息
	downloadsuite.NewMeituriSuite(firstPage, folderPath)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return err
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
