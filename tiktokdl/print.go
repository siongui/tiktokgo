package tiktokdl

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/siongui/tiktokgo"
)

var cc = color.New(color.FgCyan)
var rc = color.New(color.FgRed)

func PrintTiktokUser(user tiktokgo.TiktokUser) {
	fmt.Print("username: ")
	cc.Println(user.Nickname)

	fmt.Print("id: ")
	cc.Println(user.Id)

	if user.PrivateAccount {
		cc.Println("private account")
	} else {
		cc.Println("public account")
	}

	fmt.Print("avatar larger url: ")
	rc.Print(user.AvatarLarger)
}