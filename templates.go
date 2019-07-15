// https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
package main

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

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
	return tmpl.ExecuteTemplate(w, "layout", data) // layout -> defined in each layout template
}

// Load templates on program initialisation
func ReloadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layouts, err := filepath.Glob("templates/layouts/*.html")
	if err != nil {
		err = errors.Wrap(err, "template")
		log.Fatal(err)
	}

	pages, err := filepath.Glob("templates/pages/*.html")
	if err != nil {
		err = errors.Wrap(err, "template")
		log.Fatal(err)
	}

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
			// todo - crawl_folder func call if include is folder
			layoutBase := filepath.Base(layout)
			layoutShort := layoutBase[0:strings.LastIndex(layoutBase, ".")]
			pageBase := filepath.Base(page)
			pageShort := pageBase[0:strings.LastIndex(pageBase, ".")]
			templates[layoutShort+":"+pageShort] = template.Must(template.New(pageShort).ParseFiles(files...))
		}
	}
	log.Debugf("%v", templates)
}
