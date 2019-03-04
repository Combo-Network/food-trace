package main

import (
	"net/http"
	"fmt"

	"github.com/Combo-Network/food-trace/service"
	"github.com/Combo-Network/food-trace/web/controller"
)

func main(){

	serviceSetup := service.ServiceSetup{
	}

	app := controller.Application{
		Setup: &serviceSetup,
	}

	fs:= http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", app.LoginView)
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/queryPage", app.QueryPage)
	//http.HandleFunc("/query", app.FindByID)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Web start failed: %v", err)
	}
}