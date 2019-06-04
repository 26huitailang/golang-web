package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"syscall"
	"time"

	"github.com/26huitailang/golang-web/downloadsuite/suite"
	"github.com/feixiao/httpprof"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julienschmidt/httprouter"
)

var config *Configuration
var DB *gorm.DB

type Configuration struct {
	BasePath string `json:"base_path"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
}

// 初始化文件结构
func init() {
	initConfiguration()
	// DB 小心:= 覆盖了声明的全局变量
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Panicf("DB connect error: %s", err)
	}

	// 迁移
	DB.SingularTable(true) // 单数表名
	DB.AutoMigrate(&Theme{}, &Suite{}, &Image{})
	// sqlite 对alter table的支持有限，不支持rename column和remove column
	// err = DB.Model(&Image{}).DropColumn("IsRead").Error
}

// 加载默认配置config.json
func initConfiguration() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Error opening JSON file:", err)
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error reading JSON data:", err)
	}
	json.Unmarshal(jsonData, &config)
	initCustomConfig()
	log.Println("config:", config)
}

// 加载自定义配置，覆盖默认配置
func initCustomConfig() {
	// 文件是否存在
	file, err := os.Open("config_custom.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 读取json为map
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// var customConfig Configuration
	var customMap map[string]interface{}
	json.Unmarshal(data, &customMap)
	// log.Printf("custom config: %v", customMap)

	// 遍历config结构体，判断是否有覆盖内容
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()
	for i := 0; i < t.NumField(); i++ {
		// 比较tag
		fieldInfo := t.Field(i)
		tag := fieldInfo.Tag.Get("json")
		if value, ok := customMap[tag]; ok {
			log.Printf("tag: [%s %v] replaced by [%v]", tag, v.Field(i).Interface(), value)
			v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func themes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// var keys []string
	// for k := range Theme {
	// 	keys = append(keys, k)
	// }
	var themes []Theme
	DB.Order("name").Find(&themes)

	var t *template.Template
	t, _ = template.ParseFiles("templates/layout.html", "templates/themes.html")
	t.ExecuteTemplate(w, "layout", themes)
	//fmt.Fprintf(w, "%s\n", keys)
}

func theme(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	var t *template.Template
	t, _ = template.ParseFiles("templates/layout.html", "templates/theme.html")
	var suites []Suite
	DB.Order("name").Find(&suites)
	data := struct {
		Name   string
		Suites []Suite
	}{
		name,
		suites,
	}
	t.ExecuteTemplate(w, "layout", data)
}

func suites(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// themeName := p.ByName("name")
	suiteName := p.ByName("suite")
	var suite Suite
	var images []Image
	// for _, n := range Theme[themeName] {
	// 	if n.Name == suiteName {
	// 		data = n
	// 	}
	// }
	DB.Where("name = ?", suiteName).Find(&suite)
	DB.Model(&suite).Related(&images).Order("name")
	data := struct {
		Name   string
		Images []Image
	}{
		suite.Name,
		images,
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

func taskSuite(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// go startChild1()
	// go startChild2()
	go func() {
		s := suite.NewSuite("https://www.meituri.com/a/26718/")
		suite.DonwloadSuite(s, 5, "/Users/26huitailang/Downloads/mzitu_go", s.Title)
	}()
	fmt.Fprint(w, "task suite sent ...")
}

func taskTheme(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// url := p.ByName("url")
	var form struct {
		URL string `json:"url"`
	}
	// _ = r.ParseForm()
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err)
	}
	log.Println(form)

	go func() {
		t := suite.NewTheme(form.URL, config.BasePath)
		t.DownloadOneTheme()
		fmt.Printf("%v", t)
	}()
	fmt.Fprint(w, "task theme sent ...")
}

func initDB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// todo websocket
	log.Println("droppig table ...")
	DB.DropTableIfExists(Theme{}, Suite{}, Image{})
	log.Println("migrating table ...")
	DB.AutoMigrate(Theme{}, Suite{}, Image{})
	log.Println("start init db ...")
	InitTheme(config)
	fmt.Fprint(w, "finish init db!\n")
}

func main() {
	mux := httprouter.New()
	// profiling
	mux = httpprof.WrapRouter(mux)
	mux.GET("/", index)
	mux.GET("/hello/:name", hello)
	mux.POST("/task/suite", taskSuite)
	mux.POST("/task/theme", taskTheme)
	mux.GET("/themes", themes)
	mux.GET("/themes/:name", theme)
	mux.GET("/themes/:name/suites/:suite", suites)
	mux.POST("/initdb", initDB)
	//mux.NotFound = http.FileServer(http.Dir("/"))
	mux.ServeFiles("/image/*filepath", http.Dir(config.BasePath))

	addr := fmt.Sprintf("%s:%s", config.IP, config.Port)
	fmt.Printf("serve: http://%s\n", addr)
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	server.ListenAndServe()
}
