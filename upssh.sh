#!/bin/sh

# path为绝对路径
/path/upssh

echo "bye"

#程序退出时，同时退出堡垒机的linux服务器，防止用户恶意修改，部署时记得把注释打开
#logout