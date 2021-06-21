package view

import (
	"fmt"

	"chatClient/handle"
	"chatClient/module"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var ViewCtrl  module.ViewCtrlModule
func SetUpView() fyne.Window{

	myApp := app.New()
	myWin := myApp.NewWindow("CHAT")
	myWin.Resize(fyne.Size{Width: 500, Height: 500})


	ViewCtrl.NameEntry = widget.NewEntry()
	ViewCtrl.NameEntry.SetPlaceHolder("input name")
	ViewCtrl.NameEntry.OnChanged = func(content string) {
		handle.SetName(ViewCtrl.NameEntry.Text)
		fmt.Println("name:", ViewCtrl.NameEntry.Text, "entered")
	}
	nameBox := widget.NewHBox(widget.NewLabel("Name"), layout.NewSpacer(), ViewCtrl.NameEntry)


	ViewCtrl.SerEntry = widget.NewEntry()
	ViewCtrl.SerEntry.SetPlaceHolder("input  server        ")
	ViewCtrl.SerEntry.OnChanged = func(content string) {
		handle.SetServer(ViewCtrl.SerEntry.Text)
		fmt.Println("name:", ViewCtrl.SerEntry.Text, "entered")
	}

	ViewCtrl.ServerLab = widget.NewLabel( "server")
	ViewCtrl.UserList = widget.NewLabel("userList")
	ViewCtrl.StatuLable = widget.NewLabel("status:NO")



	ViewCtrl.SLine1 = widget.NewSeparator()				///分割线1
	ViewCtrl.SLine1.Resize(fyne.Size{Width: 500, Height: 1})
	ViewCtrl.SLine2 = widget.NewSeparator()				///分割线2
	ViewCtrl.SLine2.Resize(fyne.Size{Width: 500, Height: 1})
	ViewCtrl.SLine3 = widget.NewSeparator()				///分割线3
	ViewCtrl.SLine3.Resize(fyne.Size{Width: 500, Height: 1})


	ViewCtrl.UserChat  = widget.NewLabel("=============chat=============")

	sU := widget.NewVScrollContainer(ViewCtrl.UserChat)
	go UpdataUserChat()
	go UpdatuserList()

	ViewCtrl.Sendtext = widget.NewEntry()
	ViewCtrl.Sendtext.SetPlaceHolder("Input")
	ViewCtrl.ConBtn = widget.NewButton("con", func() {///conBtn
		if handle.Username == "" || handle.Server == ""{
			return
		}
		if handle.HandleSetupHttp() != "" {
			return
		}
		ViewCtrl.StatuLable.SetText("status:OK")
	})
	ViewCtrl.ConBtn.Resize(fyne.Size{Width: 50, Height: 100})
	ViewCtrl.DisConBtn = widget.NewButton("disCon", func() {///disConBtn
		if ViewCtrl.StatuLable.Text != "status:OK" {
			return
		}
		if handle.HandleDisCon() != ""{
			return
		}
		ViewCtrl.StatuLable.SetText("status:NO")
	})
	ViewCtrl.DisConBtn.Resize(fyne.Size{Width: 50, Height: 100})
	ViewCtrl.SendBtn = widget.NewButton("send", func() {///sendBtn
		if ViewCtrl.Sendtext.Text == "" {
			return
		}
		if handle.HandSendMsg(ViewCtrl.Sendtext.Text) != ""{
			return
		}
		ViewCtrl.Sendtext.SetText("")
	})

	ViewCtrl.SendBtn.Resize(fyne.Size{Width: 50, Height: 100})
	content := container.NewBorder(
		container.NewVBox(container.NewHBox(nameBox,ViewCtrl.NameEntry),container.NewHBox(ViewCtrl.ServerLab,ViewCtrl.SerEntry,ViewCtrl.ConBtn,ViewCtrl.DisConBtn,ViewCtrl.StatuLable),ViewCtrl.SLine2),
		container.NewVBox(ViewCtrl.SLine3,container.NewHBox(ViewCtrl.Sendtext,ViewCtrl.SendBtn)),
		container.NewHBox(ViewCtrl.UserList,ViewCtrl.SLine1,sU),
		nil)

	myWin.SetContent(content)
	return myWin
}

func UpdataUserChat() {
	for {
		select {
		case <- handle.HChan.UserChatChan:
			temp := ""
			temp = ViewCtrl.UserChat.Text + "\n" + handle.HChan.UserChatMsg
			ViewCtrl.UserChat.SetText(temp)
		default:
		}
	}
}

func UpdatuserList()  {
	for {
		select {
		case <- handle.HChan.UserListChan:
			ViewCtrl.UserList.SetText("")
			ViewCtrl.UserList.SetText(handle.HChan.UserListMsg)
		default:
		}
	}
}