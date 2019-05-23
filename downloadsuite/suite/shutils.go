package suite

// 流水线进阶版，fan-out fan-in
// 多个goroutine去读取生产数据，这个是fan-out
// 一个goroutine去获取上一步的多个输出channel，这个是fan-in

import (
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// 生产者，将1-n页 fan-out
func producer(pageUrls ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out) // 发送方关闭通道，消费方拿完后通道被回收
		for _, n := range pageUrls {
			out <- n
		}
	}()
	return out
}

// 收集结果，所有的img srcs
func merge(channels ...<-chan string) <-chan string {
	out := make(chan string, 10)
	var wg sync.WaitGroup // 群同步
	collect := func(in <-chan string) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(channels))
	// fan-in
	for _, c := range channels {
		go collect(c)
	}

	// 直接等待是死锁，因为merge写了out，main没有读，数据流入但不流出
	// 放到一个goroutine里面等待并关闭out 通道，这样不阻塞main读取
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// IsFileOrFolderExists 判断文件或文件夹是否存在
func IsFileOrFolderExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断path是不是文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断path是不是文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 获取每页的Html内容
func getURLContent(url string) []byte {
	resp, _ := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	return body
}

// 获取每页的Html内容
func getPageContent(url string) string {
	body := getURLContent(url)
	return string(body)
}
