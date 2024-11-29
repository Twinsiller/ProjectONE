package main

import (
	"ProjectONE/cmd"
	"fmt"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Printf("Запуск программы не сработал!!!\n%v", err)
	}
}
