package module

import (
	"fyne.io/fyne/widget"
)
type  ViewCtrlModule struct {
	NameEntry *widget.Entry
	SerEntry *widget.Entry
	Sendtext *widget.Entry


	ServerLab *widget.Label
	UserList *widget.Label
	StatuLable *widget.Label
	UserChat *widget.Label

	SLine1 *widget.Separator
	SLine2 *widget.Separator
	SLine3 *widget.Separator

	ConBtn *widget.Button
	DisConBtn *widget.Button
	SendBtn *widget.Button
}
