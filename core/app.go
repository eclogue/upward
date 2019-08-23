package core

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
	"strconv"
	"strings"
)
//type Server struct {
//	Id       int         `json:"id"`
//	Name     string      `json:"name"`
//	Host     string      `json:"host"`
//	Desc     string      `json:"desc"`
//}

type App struct {
	Servers map[string]Server
}

func (app *App)Init(){
	//输出服务器列表
	app.showServers()

	// 监听输入
	//input, isGlobal := app.checkInput()
	data := make(chan []interface{})
	StartInput := make(chan bool)
	go app.checkInput(data,StartInput)

	for {
		select {
		case dataList := <-data:
			//fmt.Println(dataList)

			input := dataList[0].(string)
			isGlobal := dataList[1].(bool)

			if isGlobal {
				end := app.handleGlobalCmd(input)
				if end {
					return
				}
			} else {
				server := app.Servers[input]
				Printer.Infoln("你选择了", server.Name)
				Log.Category("app").Info("select server", server.Name)

				StopInput = true
				server.Connect()
			}

			Printer.Infoln("按任意键回到首页")
			fmt.Scanln()

			//fmt.Println(1)
			app.showServers()
			StopInput = false
			StartInput <- true
			//fmt.Println(2)
		}
	}

	//if isGlobal {
	//	end := app.handleGlobalCmd(input)
	//	if end {
	//		return
	//	}
	//} else {
	//	server := app.Servers[input]
	//	Printer.Infoln("你选择了", server.Name)
	//	Log.Category("app").Info("select server", server.Name)
	//	server.Connect()
	//}
}

func (app *App) GetServers(username string){
	servers := GetUserHosts(username)
	app.Servers = make(map[string]Server)
	for index,value := range servers {
		app.Servers[strconv.Itoa(index+1)] = value
	}
}

// 打印列表
func (app *App) showServers() {

	maxlen := 50.0
	app.formatSeparator(" 欢迎使用 Upward SSH ", "=", maxlen)

	//fmt.Println(app.Servers["1"])
	length := len(app.Servers)
	for i:=1;i<=length ;i++  {
		Printer.Logln(app.recordServer(strconv.Itoa(i), app.Servers[strconv.Itoa(i)]))
	}

	//for i, server := range app.Servers {
	//	Printer.Logln(app.recordServer(i, server))
	//}

	app.formatSeparator("", "=", maxlen)
	Printer.Logln("", "[exit]\t退出")
	app.formatSeparator("", "=", maxlen)
	Printer.Info("请输入序号或操作: ")
}

func (app *App) formatSeparator(title string, c string, maxlength float64) {

	charslen := int((maxlength - ZhLen(title)) / 2)
	chars := ""
	for i := 0; i < charslen; i ++ {
		chars += c
	}

	Printer.Infoln(chars + title + chars)
}

func (app *App) recordServer(flag string, server Server) string {
		return " [" + flag + "]" + "\t" + server.Name + "\t"+ " [" + server.Host + "]" + "\t"+ " (" + server.Desc + ")"
}

// 检查输入
//func (app *App) checkInput() (string, bool) {
func (app *App) checkInput(data chan []interface{},startInput chan bool) {
	for {
		var flag string
		if StopInput {
			select {
			case <-startInput:
				StopInput = false
				//fmt.Println(3)
				continue
			}
		}else{
			//fmt.Println(4)

			fmt.Scanln(&flag)

			if app.isGlobalInput(flag) {
				data <- []interface{}{flag, true}
				continue
				//return flag, true
			}

			if _, ok := app.Servers[flag]; !ok {
				Printer.Errorln("输入有误，请重新输入")
			} else {
				data <- []interface{}{flag, false}
				StopInput = true
				continue
				//return flag, false
			}
		}
	}

	panic(errors.New("输入有误"))
}

// 判断是否全局输入
func (app *App) isGlobalInput(flag string) bool {
	switch flag {
	case "exit":
		return true
	default:
		return false
	}
}

func (app *App) handleGlobalCmd(cmd string) bool {
	switch strings.ToLower(cmd) {
	case "exit":
		return app.runCmd(cmd)
	default:
		Printer.Errorln("指令无效")
		return false
	}
}

func (app *App) runCmd(command string) bool {
	//command = "logout"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash","-c",command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	//fmt.Println("stdout: ",stdout.String())
	//fmt.Println("stderr: ",stderr.String())
	if err != nil {
		//fmt.Println(err.Error())
		Printer.Errorln("指令无效")
		return false
	}
	return true
}