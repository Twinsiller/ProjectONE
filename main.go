package main

import "ProjectONE/cmd"

func main() {
	if cmd.Run() != nil {
		print("Запуск программы не сработал!!!")
	}
}
