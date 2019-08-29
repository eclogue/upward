package core

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"upward/config"
)

func CheckUser(username string, host string, pwd string) (bool, error) {
	type data struct {
		Token    string
		Username string
	}

	result := struct {
		Code    int
		Message string
		Data    data
	}{}
	resp, err := http.PostForm(config.Domain+"/fortress/login", url.Values{"username": {username}, "host": {host}, "pwd": {pwd}})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, err
	}

	if result.Code == 0 {
		JwtToken = result.Data.Token
		return true, nil
	} else {
		return false, errors.New(result.Message)
	}
}

func GetUserHosts(username string) (list []Server) {
	param := map[string]interface{}{"username": username}
	body, err := HttpPost(config.Domain+"/fortress/hosts", param, nil, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	type data struct {
		_id              string
		Ansible_ssh_host string
		Ansible_ssh_user string
		Hostname         string
	}

	result := struct {
		Code    int
		Message string
		Data    []data
	}{}

	//fmt.Println(string(body))
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(result.Data)
	for k, v := range result.Data {
		server := Server{Id: k + 1, Name: v.Ansible_ssh_user, Host: v.Ansible_ssh_host, Desc: v.Hostname}
		list = append(list, server)
	}
	return list
}
