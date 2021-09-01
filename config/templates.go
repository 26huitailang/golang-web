// https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
package config

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	rice "github.com/GeertJohan/go.rice"

	"github.com/labstack/echo/v4"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//go:generate rice embed-go
// folder structure

// views/layouts/base.hbs
// -----------------------
// [[define "layout"]] [[block "content" . ]] default content [[end]] [[end]]

// views/pages/contact.hbs
// -----------------------
// [[define "content"]] contact me [[block "ad" .]] [[end]] [[end]]

// views/pages/home.hbs
// -----------------------
// [[define "content"]] welcome [[block "ad" .]] [[end]]  [[end]]

// views/shared/ad.hbs
// [[define "ad"]] buy stuff [[end]]

// init:
// var EchoTemplate = &Template{}
// echo_inst.SetRenderer(EchoTemplate)
// ReloadTemplates onStart and while developing viewstuff in handler for example
//

// useage in handler
// func H_Start(c *echo.Context) error {
//	return c.Render(200, "base:home", nil)
// }

var templates map[string]*template.Template

type Template struct {
	templates *template.Template // this will never be used
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := templates[name]
	if !ok {
		err := errors.Errorf("template name not found -> %s template", name)
		log.Errorf("templates: name not found -> %s template", name)
		return err
	}
	err := tmpl.ExecuteTemplate(w, "layout", data) // layout -> defined in each layout template
	if err != nil {
		log.Panicf("Render: %v", err)
	}
	return err
}

// Load templates on program initialisation
func ReloadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templateBox, err := rice.FindBox("../templates")
	if err != nil {
		log.Fatal("find box: ", err)
	}
	// log.Debug("templateBox ", templateBox.Name())

	layouts := make([]string, 0)
	pages := make([]string, 0)
	templateBox.Walk("", func(path string, info os.FileInfo, err error) error {
		//log.Info(path)
		if !info.IsDir() {
			if strings.Contains(path, "layouts") {
				layouts = append(layouts, path)
			} else if strings.Contains(path, "pages") {
				pages = append(pages, path)
			}
		}
		if err != nil {
			return err
		}
		return nil
	})
	// log.Info("layouts ", layouts)
	//layouts, err := filepath.Glob("templates/layouts/*.html")
	//if err != nil {
	//	err = errors.Wrap(err, "template")
	//	log.Fatal(err)
	//}
	//
	//pages, err := filepath.Glob("templates/pages/*.html")
	//if err != nil {
	//	err = errors.Wrap(err, "template")
	//	log.Fatal(err)
	//}

	globalShared, err := filepath.Glob("views/shared/*.html")
	if err != nil {
		err = errors.Wrap(err, "template")
		log.Fatal(err)
	}

	// todo: var crawl_folder = func(){}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		for _, page := range pages {
			files := append(globalShared, layout, page)
			// log.Info("files", files)
			// todo - crawl_folder func call if include is folder
			layoutBase := filepath.Base(layout)
			layoutShort := layoutBase[0:strings.LastIndex(layoutBase, ".")]
			pageBase := filepath.Base(page)
			pageShort := pageBase[0:strings.LastIndex(pageBase, ".")]
			//pageShort := pageBase
			content := ""
			for _, file := range files {
				file = strings.Replace(file, "\\", "/", -1)
				tmp, err := templateBox.String(file)
				if err != nil {
					log.Fatal(err)
				}
				content = content + tmp
			}
			templates[layoutShort+":"+pageShort] = template.Must(template.New(pageShort).Parse(content))
		}
	}
	log.Debugf("templates %v", templates)
}
