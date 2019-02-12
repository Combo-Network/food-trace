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
	http.HandleFunc("/", app.LoginView)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Web start failed: %v", err)
	}
}