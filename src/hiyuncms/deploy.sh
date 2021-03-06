export GOPATH=/Users/liusheng/GoglandProjects/martini
echo "1、------------------------开始编译 Main.go------------------------------"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build Main.go
echo "1、---------------------------------------------------------------------"

echo "2、---------------------------停止服务器上 main---------------------------"
ssh -tt 180.76.187.132 << eeooff
    ps aux | grep Main | grep -v grep | awk '{print "kill ",\$2|"bash"}'
    exit
eeooff
echo "2、---------------------------------------------------------------------"

echo "3、----------------------------开始复制配置文件----------------------------"
scp ./config.ini 180.76.187.132:/home/webroot/conf/
echo "3、---------------------------------------------------------------------"

echo "4、----------------------------上传 main--------------------------------"
scp ./Main 180.76.187.132:/home
echo "4、--------------------------------------------------------------------"

echo "5、---------------------------启动服务器上 main---------------------------"
ssh -tt 180.76.187.132  << eeooff
   cd /home
   nohup ./Main >out.file 2>&1 &
   exit
eeooff
echo "5、---------------------------------------------------------------------"

echo "6、----------------------------开始复制模板文件----------------------------"
scp -r ./webroot/templates 180.76.187.132:/home/webroot
echo "6、---------------------------------------------------------------------"

echo "6、----------------------------开始复制静态文件----------------------------"
#scp -r ./webroot/static 180.76.187.132:/home/webroot
echo "6、---------------------------------------------------------------------"