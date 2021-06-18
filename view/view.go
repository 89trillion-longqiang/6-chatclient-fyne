package view

import (
	"fmt"

	"chatClient/handle"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type View struct {
	UserList *widget.Label
	UserChat * widget.Label
}

var ViewLable View


func SetUpView() fyne.Window{

	myApp := app.New()
	myWin := myApp.NewWindow("CHAT")
	myWin.Resize(fyne.Size{Width: 500, Height: 500})

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("input name")
	nameEntry.OnChanged = func(content string) {
		handle.SetName(nameEntry.Text)
		fmt.Println("name:", nameEntry.Text, "entered")
	}
	nameBox := widget.NewHBox(widget.NewLabel("Name"), layout.NewSpacer(), nameEntry)


	serEntry := widget.NewEntry()
	serEntry.SetPlaceHolder("input  server        ")
	serEntry.OnChanged = func(content string) {
		handle.SetServer(serEntry.Text)
		fmt.Println("name:", serEntry.Text, "entered")
	}

	serverLab := widget.Label{Text: "server"}
	userList := widget.NewLabel("userList")
	statuLable := widget.NewLabel("status:NO")



	sLine1 := widget.NewSeparator()				///分割线1
	sLine1.Resize(fyne.Size{Width: 500, Height: 1})
	sLine2 := widget.NewSeparator()				///分割线2
	sLine2.Resize(fyne.Size{Width: 500, Height: 1})
	sLine3 := widget.NewSeparator()				///分割线3
	sLine3.Resize(fyne.Size{Width: 500, Height: 1})


	userChat  := widget.NewLabel("=============chat=============")

	ViewLable.UserList = userList
	ViewLable.UserChat = userChat
	sU := widget.NewVScrollContainer(userChat)
	go UpdataUserChat()
	go UpdatuserList()

	sendtext := widget.NewEntry()
	sendtext.SetPlaceHolder("Input")
	conBtn := widget.NewButton("con", func() {///conBtn
		if handle.Username == "" || handle.Server == ""{
			return
		}
		if handle.HandleSetupHttp() != "" {
			return
		}
		statuLable.SetText("status:OK")
	})
	conBtn.Resize(fyne.Size{Width: 50, Height: 100})
	disConBtn := widget.NewButton("disCon", func() {///disConBtn
		if statuLable.Text != "status:OK" {
			return
		}
		if handle.HandleDisCon() != ""{
			return
		}
		statuLable.SetText("status:NO")
	})
	disConBtn.Resize(fyne.Size{Width: 50, Height: 100})
	sendBtn := widget.NewButton("send", func() {///sendBtn
		if sendtext.Text == "" {
			return
		}
		if handle.HandSendMsg(sendtext.Text) != ""{
			return
		}
		sendtext.SetText("")
	})

	sendBtn.Resize(fyne.Size{Width: 50, Height: 100})
	content := container.NewBorder(
		container.NewVBox(container.NewHBox(nameBox,nameEntry),container.NewHBox(&serverLab,serEntry,conBtn,disConBtn,statuLable),sLine2),
		container.NewVBox(sLine3,container.NewHBox(sendtext,sendBtn)),
		container.NewHBox(userList,sLine1,sU),
		nil)

	myWin.SetContent(content)
	return myWin
}

func UpdataUserChat() {
	for {
		select {
		case <- handle.HChan.UserChatChan:
			temp := ""
			temp = ViewLable.UserChat.Text + "\n" + handle.HChan.UserChatMsg
			ViewLable.UserChat.SetText(temp)
		default:
		}
	}
}

func UpdatuserList()  {
	for {
		select {
		case <- handle.HChan.UserListChan:
			ViewLable.UserList.SetText("")
			ViewLable.UserList.SetText(handle.HChan.UserListMsg)
		default:
		}
	}
}