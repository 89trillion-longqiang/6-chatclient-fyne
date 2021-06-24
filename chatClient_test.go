package main

import (
	"chatClient/controller"
	"testing"
)



func Test_SetName(t *testing.T) {

	controller.SetName("username")
	if controller.Username != "username"{
		t.Errorf("Test_SetName error ")
	}else {
		t.Log("Test_SetName pass ")
	}

}

func Test_SetServer(t *testing.T){
	controller.SetServer("server")
	if controller.Server != "server"{
		t.Errorf("Test_SetServer error ")
	}else {
		t.Log("Test_SetServer pass ")
	}
}