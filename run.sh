#!/bin/bash
node="allstats"
killall -9 bee > /dev/null
source /root/.profile
source /root/.bashrc
GOPATH="/root/go"
Home=/root 
if [ -z $GOPATH ];then
    echo "GOPATH not set!"
    exit 1
fi
if [ -f $GOPATH/bin/bee ];then
    go get -u github.com/astaxie/beego
    go get -u github.com/beego/bee
fi
cd $GOPATH/src/$node
$GOPATH/bin/godep get
$GOPATH/bin/bee run
