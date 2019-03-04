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

var cuser User

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")

	var flag bool
	for _, user := range users {
		if user.LoginName == loginName && user.Password == password {
			cuser = user
			flag = true
			break
		}
	}

	data := &struct {
		CurrentUser User
		Flag bool
	}{
		CurrentUser:cuser,
		Flag:false,
	}

	if flag {
		ShowView(w, r, "index.html", data)
	}else{
		data.Flag = true
		data.CurrentUser.LoginName = loginName
		ShowView(w, r, "login.html", data)
	}
}

func (app *Application) QueryPage(w http.ResponseWriter, r *http.Request)  {
	data := &struct {
		CurrentUser User
		Msg string
		Flag bool
	}{
		CurrentUser:cuser,
		Msg:"",
		Flag:false,
	}
	ShowView(w, r, "query.html", data)
}
