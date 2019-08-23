package core

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//生成client 通过client的Do方法执行
func HttpPost(url string,params map[string]interface{},header map[string]string,timeout time.Duration)([]byte,error){
	
	/*     生成client,参数默认;
	*    这个结构体有四个属性
	*    Transport(RoundTrepper);指定执行的独立、单次http请求的机制
	*    CheckRedirect(func(req *Request, via []*Request)err):指定处理重定向的策略,如果不为nil,客户端在执行重定向之前调用本函数字段.如果CheckRedirect 返回一个错误，本类型中的get方法不会发送请求,如果CheckRedirect为nil,就会采用默认策略:连续请求10次后停止；
	＊    jar(CookieJar):jar指定cookie管理器,若果为nil请求中不会发送cookie,回复中的cookie会被忽略
	＊    TimeOut(time.Duration):指定本类型请求的时间限制，为0表示不设置超时
	*/
	//client := &http.Client{}    这里初始化了一个默认的client，并没有配置一些请求的设置

	//可以通过client中transport的Dail函数,在自定义Dail函数里面设置建立连接超时时长和发送接受数据超时
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second * timeout)    //设置建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * timeout))    //设置发送接受数据超时
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * timeout,
		},
	}

	stringParams := ""
	for k,v := range params{
		stringParams += k
		stringParams += "="
		switch v.(type) {
		case string:
			stringParams += v.(string)
			break
		case int:
			stringParams += strconv.Itoa(v.(int))
			break
		case float32:
			stringParams += strconv.FormatFloat(float64(v.(float32)),'f',6,64)
			break
		case float64:
			stringParams += strconv.FormatFloat(v.(float64),'f',6,64)
			break
		default:
			return nil,errors.New("param error")
		}
		stringParams += "&"
	}
	stringParams = strings.TrimRight(stringParams,"&")
	reqest,err := http.NewRequest("POST",url,strings.NewReader(stringParams))    //提交请求;用指定的方法，网址，可选的主体放回一个新的*Request


	for key,value := range header{
		reqest.Header.Set(key,value)
	}
	reqest.Header.Set("Content-Type","application/x-www-form-urlencoded")
	if err != nil {
		//panic(err)
		return nil,err
	}
	response,err := client.Do(reqest)    //前面预处理一些参数，状态，Do执行发送；处理返回结果;Do:发送请求,
	if err != nil {
		return nil,err
	}
	defer response.Body.Close()
	//stdout := os.Stdout            //将结果定位到标准输出，也可以直接打印出来，或定位到其他地方进行相应处理
	//_,err = io.Copy(stdout,response.Body)    //将第二个参数拷贝到第一个参数，直到第二参数到达EOF或发生错误，返回拷贝的字节喝遇到的第一个错误.
	status := response.StatusCode        //获取返回状态码，正常是200
	if status == 200 {
		body,err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil,err
		}
		return body,err
	}else{
		return nil,errors.New("request failed")
	}

	//log.Println(string(body))
	//log.Println(status)

}