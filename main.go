package main

import (
	"fmt"
	"go_todolist/db"
	"go_todolist/routes"
)

func main() {
	fmt.Println("Hello world")
	db.Init()
	router := routes.SetUpRouter()
	router.Run(":8080")
}
