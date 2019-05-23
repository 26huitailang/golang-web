package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/26huitailang/golang-web/downloadsuite/suite"
	"github.com/feixiao/httpprof"
	"github.com/julienschmidt/httprouter"
)

var config Configuration
var Theme map[string][]*Suite

type Configuration struct {
	BasePath string `json:"base_path"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
}

type Suite struct {
	Name   string
	Images []Image
}

type Image struct {
	Path string
}

// 初始化文件结构
func init() {
	Theme = make(map[string][]*Suite)
	InitConfiguration()
	InitTheme(&config)
}

func InitConfiguration() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	json.Unmarshal(jsonData, &config)
	fmt.Println(config)
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func themes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var keys []string
	//a := map[string]*Suite{"a": &Suite{}, "b": &Suite{}}
	for k := range Theme {
		keys = append(keys, k)
	}
	var t *template.Template
	t, _ = template.ParseFiles("templates/layout.html", "templates/themes.html")
	t.ExecuteTemplate(w, "layout", keys)
	//fmt.Fprintf(w, "%s\n", keys)
}

func theme(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	var t *template.Template
	t, _ = template.ParseFiles("templates/layout.html", "templates/theme.html")
	data := struct {
		Name   string
		Suites []*Suite
	}{
		name, Theme[name],
	}
	t.ExecuteTemplate(w, "layout", data)
}

func suites(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	themeName := p.ByName("name")
	suiteName := p.ByName("suite")
	var data *Suite
	for _, n := range Theme[themeName] {
		if n.Name == suiteName {
			data = n
		}
	}
	var t *template.Template
	t, _ = template.ParseFiles("templates/layout.html", "templates/suite.html")
	t.ExecuteTemplate(w, "layout", data)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Redirect(w, r, "/themes", 301)
}

func startChild1() {
	cmd := exec.Command("/bin/sh", "-c", "sleep 1000")
	time.AfterFunc(10*time.Second, func() {
		fmt.Println("PID1=", cmd.Process.Pid)
		syscall.Kill(-cmd.Process.Pid, syscall.SIGQUIT)
		fmt.Println("killed")
	})
	fmt.Println("begin run")
	cmd.Run()
}
func startChild2() {
	for index := 0; index < 10; index++ {
		time.Sleep(1 * time.Second)
		fmt.Println(index)
	}
}

func task1(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	go startChild1()
	go startChild2()
	go func() {
		s := suite.NewSuite("https://www.meituri.com/a/26718/")
		suite.DonwloadSuite(s, 5, "/Users/26huitailang/Downloads/mzitu_go", s.Title)
	}()
	fmt.Fprint(w, "task running")
}

func main() {
	mux := httprouter.New()
	// profiling
	mux = httpprof.WrapRouter(mux)
	mux.GET("/", index)
	mux.GET("/hello/:name", hello)
	mux.GET("/task", task1)
	mux.GET("/themes", themes)
	mux.GET("/themes/:name", theme)
	mux.GET("/themes/:name/suites/:suite", suites)
	//mux.NotFound = http.FileServer(http.Dir("/"))
	mux.ServeFiles("/image/*filepath", http.Dir(config.BasePath))

	addr := fmt.Sprintf("%s:%s", config.IP, config.Port)
	fmt.Printf("serve: http://%s", addr)
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	server.ListenAndServe()
}
