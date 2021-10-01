package main

import (
	"fmt"
)
// App - the struct which contains things like, pointers to the database
// connections
type App struct {}

func(app *App) Run() error{
	fmt.Println("Setting Up Our APP")
	return nil

}

func main(){
	fmt.Println("Go REST API Project")
	app:=App{}
	if err:=app.Run(); err!=nil{
		fmt.Println("Error starting up or REST API")
		fmt.Println(err)
	}
}
