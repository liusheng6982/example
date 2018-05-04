#!/bin/sh
#

compile()
{
  export GOPATH=/Users/liusheng/GoglandProjects/martini
  echo "1、------------------------开始编译 Main.go------------------------------"
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o demo_cms Main.go
  echo "1、----------------------------------------------------------------------"
}

stopServer()
{
  echo "2、---------------------------停止服务器上 main---------------------------"
  ssh -tt 180.76.187.132 << eeooff
        ps aux | grep demo_cms | grep -v grep | awk '{print "kill ",\$2|"bash"}'
        exit
  eeooff
  echo "2、---------------------------------------------------------------------"
}

copyConf()
{
    echo "3、----------------------------开始复制配置文件----------------------------"
    scp ./config.ini 180.76.187.132:/root/demo/webroot/conf/
    echo "3、---------------------------------------------------------------------"
}

copyMain()
{
    echo "4、----------------------------上传 main--------------------------------"
    scp ./demo_cms 180.76.187.132:/root/demo/
    echo "4、--------------------------------------------------------------------"
}

startServer()
{
  echo "5、---------------------------启动服务器上 main---------------------------"
  ssh -tt 180.76.187.132  << eeooff
    cd /root/demo
    nohup ./demo_cms >out.file 2>&1 &
    exit
eeooff
  echo "5、---------------------------------------------------------------------"
}

copyTemplate()
{
  echo "6、----------------------------开始复制模板文件----------------------------"
  scp -r ./webroot/templates 180.76.187.132:/root/demo/webroot
  echo "6、---------------------------------------------------------------------"
}

copyStatic()
{
  echo "7、----------------------------开始复制静态文件----------------------------"
  scp -r ./webroot/static 180.76.187.132:/root/demo/webroot
  echo "7、---------------------------------------------------------------------"
}


#compile
#stopServer
#copyConf
#copyMain
#startServer
copyTemplate
#copyStatic

