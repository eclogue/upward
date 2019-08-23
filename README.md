#upward_ssh
一个远程连接跳板机的ssh客户端

## 版本说明

## 功能说明
- 远程连接跳板机

## 下载

## 安装
- 下载编译好的二进制包upssh，放在指目录下，如`~/upssh`或`/usr/local/upssh`

## 使用说明
- 在堡垒机某用户目录下的.bash_profile文件后面，添加如下：／path／upssh.sh，即可
 
## 编译
go build upssh.go

## 依赖包
- golang.org/x/crypto/ssh