package main

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"golang_web/downloadsuite"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type myMessageHandler struct{}

func processMessage(body []byte) error {
	imageInfo := new(downloadsuite.ImageInfo)
	if err := json.Unmarshal(body, imageInfo); err != nil {
		return err
	}

	isFolderExist := downloadsuite.IsFileOrFolderExists(imageInfo.Path)
	if !isFolderExist {
		fmt.Println("创建文件夹: ", imageInfo.Path)
		err := os.MkdirAll(imageInfo.Path, os.ModePerm)
		if err != nil {
			log.Printf("%s", err)
			return err
		}
	}

	name := path.Join(imageInfo.Path, imageInfo.Name)
	if downloadsuite.IsFileOrFolderExists(name) {
		fmt.Println("已存在: ", name)
		return nil
	}
	content := downloadsuite.GetURLContent(imageInfo.URL)
	if err := ioutil.WriteFile(name, content, 0644); err != nil {
		fmt.Println("failed: ", imageInfo.URL)
		return err
	}
	log.Printf("donwloaded successfully: %s %s", imageInfo.Name, imageInfo.URL)

	return nil
}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}
	err := processMessage(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return err
}

func consume() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("mtr_image", "channel", config)
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
