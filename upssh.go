package main

import (
	"fmt"
	"os/user"
	"upward/config"
	"upward/core"
)

func main() {
	config.Init()

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