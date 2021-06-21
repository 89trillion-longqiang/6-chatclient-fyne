package main

import (
	"chatClient/handle"
	"testing"
)



func Test_SetName(t *testing.T) {

	handle.SetName("username")
	if handle.Username != "username"{
		t.Errorf("Test_SetName error ")
	}else {
		t.Log("Test_SetName pass ")
	}

}

func Test_SetServer(t *testing.T){
	handle.SetServer("server")
	if handle.Server != "server"{
		t.Errorf("Test_SetServer error ")
	}else {
		t.Log("Test_SetServer pass ")
	}
}