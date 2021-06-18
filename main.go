package main

import (
	"log"
	"os"

	"chatClient/view"
)

func main() {

	os.Mkdir("log", 0777)
	file, err := os.OpenFile("log/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	myWin := view.SetUpView()
	myWin.ShowAndRun()
}
