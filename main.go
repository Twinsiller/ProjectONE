package main

import (
	"ProjectONE/cmd"
	"fmt"
)

//@title ProjectONE
//@version 1.0
//@description Project for project...

//@host localhost:8080
//@BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := cmd.Run(); err != nil {
		fmt.Printf("Запуск программы не сработал!!!\n%v", err)
	}
}
