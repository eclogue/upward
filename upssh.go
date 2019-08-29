package main

import (
	"fmt"
	"os/user"
	"upward/core"
)

func main() {
	u,_ := user.Current()
	username := u.Username

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			core.Log.Category("upssh").Error("recover", err)
		}
	}()

	app := new(core.App)
	app.GetServers(username)
	app.Init()
}