package module

type HandleChan struct {
	UserListChan chan int
	UserChatChan chan int
	UserListMsg string
	UserChatMsg string
}
