package main

import (
	"fmt"
	"os/user"
	"upward_ssh/core"
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

	/*if len(os.Args) == 2 {
		option := strings.Split(os.Args[1], "=")

		switch option[0] {
		case "--help":
			fallthrough
		case "-h":
			help()
			return
		case "--version":
			fallthrough
		case "-v":
			version()
			return
		}

		index := strings.Index(option[0],"@")
		if index > 0 {
			user := string([]byte(option[0])[:index])
			host := string([]byte(option[0])[index+1:])
			//fmt.Println("username: ",user)
			//fmt.Println("host: ",host)

			fmt.Printf(option[0]+"'s password: ")
			pwd,err := gopass.GetPasswd()
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println(string(pwd))
			if res,err := core.CheckUser(user,host,string(pwd)); res==false||err != nil {
				fmt.Println(err.Error())
				core.Log.Category("upssh").Error("authorization", err.Error())
				return
			}*/

	app := new(core.App)
	app.GetServers(username)
	app.Init()

		/*}else{
			fmt.Println("参数错误，详情请查看--help")
			return
		}
	}else{
		fmt.Println("参数错误，详情请查看--help")
		return
	}*/
}


// 版本信息
func version() {
	fmt.Println("Version : v1.0.0")
}

// 显示帮助信息
func help() {
	fmt.Println("一个ssh远程客户端，可一键登录远程服务器。")
	fmt.Println("参数：")
	fmt.Println("  -h, --help   ", "                      \t", "显示帮助信息。")
	fmt.Println("  -v, --version", "                      \t", "显示 upward-ssh 的版本信息。")
}