package controller

import (
	"net/http"
	"fmt"
	"path/filepath"
	"html/template"
	"github.com/Combo-Network/food-trace/service"
)

type Application struct {
	Setup *service.ServiceSetup
}

func ShowView(w http.ResponseWriter, r *http.Request, templateName string, data interface{})  {

	pagePath := filepath.Join("tpl", templateName)

	resultTemplate, err := template.ParseFiles(pagePath)
	if err != nil {
		fmt.Printf("Create template instance err: %v", err)
		return
	}

	err = resultTemplate.Execute(w, data)
	if err != nil {
		fmt.Printf("Execute template instance err: %v", err)
		return
	}

}

func (app *Application) LoginView(w http.ResponseWriter, r *http.Request)  {

	ShowView(w, r, "login.html", nil)
}