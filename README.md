## upward_ssh
- 一个跳板机的ssh服务

## 功能说明
- 远程连接跳板机

## 下载
- linux amd64版 编译好的二进制包地址：upward_ssh/bin/linux/amd64/upssh
- mac版 编译好的二进制包地址：upward_ssh/bin/mac/upssh

## 安装
- 下载编译好的二进制包upssh，放在指目录下，如`~/upssh`或`/usr/local/upssh`

## 使用说明
- 在堡垒机某用户目录下的.bash_profile文件后面，添加如下：/path/upssh.sh，即可
 
## 编译
- go build upssh.go

## 依赖包
- golang.org/x/crypto/ssh