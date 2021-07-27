#### ubuntu linux 安装和软件配置 

```shell
#######设置密码和下载安装文件 start ############
设置管理员密码：
sudo passwd
curl -o 本地文件名  链接
chmod +x 本地文件/chmo 777 本地文件
sudo chmod -R 777 temp  修改文件夹权限
mv * ../ 当前所有文件移动到上一级
mkdir test  创建文件夹
touch aa.txt  创建文件
sudo dpkg-reconfigure dash  选No
rm:
-r 就是向下递归，不管有多少级目录，一并删除
-f 就是直接强行删除，不作任何提示的意思
删除文件夹实例：
rm -rf /var/log/httpd/access 将会删除/var/log/httpd/access目录以及其下所有文件、文件夹
删除文件使用实例：
rm -f /var/log/httpd/access.log 将会强制删除/var/log/httpd/access.log这个文件
复制并提示是否要覆盖
cp -ri /home/wisn/  /home/wisn/1

#######设置密码和下载文件 end ############
```



#### 传输文件配置 

```shell
#######传输文件 start ############
scp /Users/mac/gomodproject/skydrive/doc/skydrive.sql wisn@172.17.57.196:/home/wisn/
https://www.cnblogs.com/withfeel/p/10635873.html
#######传输文件 end ############
```

#### 解压缩文件 

```shell
#######解压缩文件 start ############
unzip -o -d /home/sunny myfile.zip
zip -d myfile.zip smart.txt

删除压缩文件中smart.txt文件
zip -m myfile.zip ./rpm_info.txt
#######解压缩文件 end ############
```



#### mysql 

```shell
#######mysql start ############
是否已经安装myslq
sudo netstat -tap | grep mysql
sudo apt-get install mysql-server mysql-client
sudo mysql-uroot -p 进入后
设置root密码
        方法1： 用SET PASSWORD命令
        　　首先登录MySQL。
        格式：mysql> set password for 用户名@localhost = password('新密码');
        例子：mysql> set password for root@localhost = password('123');
        方法2：用mysqladmin
        　　格式：mysqladmin -u用户名 -p旧密码 password 新密码
        例子：mysqladmin -uroot -p123456 password 123
        方法3：用UPDATE直接编辑user表
        　　首先登录MySQL。
        mysql> use mysql;
        mysql> update user set password=password('123') where user='root' and host='localhost';
        mysql> flush privileges;

一直报语法错误用这种
  ALTER USER 'root'@'localhost' IDENTIFIED BY 'nihao@123456';
connect mysql failed Error 1698: Access denied for user 'root'@'localhost'
解决办法：ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'nihao@123456';
$ sudo mysql -u root
mysql> ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'nihao@123456';
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'nihao@123456';
$ sudo service mysql stop
$ sudo service mysql start
$ mysql -u root -p

问题：
linux  2003 - Can't connect to MySQL
解决：
update user set host='%' where user='root';
grant all privileges on *.* to 'root'@'%'
flush privileges;
https://blog.csdn.net/freezingxu/article/details/77088506


 mysql_upgrade -u root -p 

数据库备份
mysqldump -uroot -pnihao@123456 --databases  skydrive > /home/wisn/skydrive-$(date +%Y%m%d).sql
#######mysql end ############
```



#### SSH

```shell
#######SSH start ############

远程链接：
 ssh  root@localhost

 sudo apt-get install ssh
 sudo service ssh start

 检查端口是否监听
 ss -tnlp
#######SSH end ############
```





#### 自启动

```shell
#######自启动 start ############
/etc/rc.local
#!/bin/bash
#执行的命令
(
sleep 10
执行脚本  &
) &
exit 0
#######自启动 end ############
```



#### 硬盘

```shell
#######添加硬盘 start ############
 #查看更多硬盘
 fdisk -l

操作sdb
fdisk /dev/sdb
 然后，我们为这个硬盘创建分区，输入fdisk /dev/sdb，依次输入n，p，1，w，
 其中n分别表示创建一个新分区，p表示分区类型为主分区，1表示分区编号是1，w表示保存
mkfs.ext4 /dev/sdb1
下一步是格式化分区，我们输入mkfs.ext4 /dev/sdb1 （1就是上一步的分区编号）。

mkdir /data
mkdir /data，在根目录创建/data作为此分区的挂载点，
mount /dev/sdb1 /data
输mount /dev/sdb1 /data，将分区挂载到目录下，
df -h |grep sdb1
通过df -h,可以看到挂载成功

自动挂载，输vi /etc/fstab
/dev/sdb1 /data ext4 defaults 0 0

df -h
输入df -h检查一下，分区自动挂载

https://www.cnblogs.com/zishengY/p/7137671.html

査看文件的详细信息
stat demo.txt
#######添加硬盘 end ############
```





#######执行程序 start ############
ServeLocation              = 9996
	UDP_SERVER_ListenPORT      = 8997
	UDP_GroupSERVER_ListenPORT = 8998

```ini
UdpGroupserverSendport = 8999
UdpServerSendport      = 8996

USER_NAME    = "root"
PASS_WORD    = "nihao@123456"
HOST         = "localhost"
PORT         = "3306"
DATABASE     = "skydrive"
```
#### 执行可运行二进制报错

```
Syntax error: "(" unexpected（linux系统）:
	sudo dpkg-reconfigure dash
	选No
```

#######执行程序 end ############

#######执行程序 end ############

#### 基本操作

```
#######kill pid start ############
 ps -aux|grep "start.sh"
ps
kill -l pid
-l选项告诉kill命令用好像启动进程的用户已注销的方式结束进程

kill -9 PID  强大和危险的命令迫使进程在运行时突然终止
#######kill pid end ############

root@skydriver:/home/wisn/deploy# cp skydrive_release ../
cp: cannot create regular file '../skydrive_release': Text file busy
root@skydriver:/home/wisn# fuser skydrive_release 
/home/wisn/skydrive_release:   690e
root@skydriver:/home/wisn# kill -9 690e
bash: kill: 690e: arguments must be process or job IDs
root@skydriver:/home/wisn# kill -9 690







#######查看文件大小 start ############
只显示直接子目录文件及文件夹大小统计值：
du -h --max-depth=1 thumbnail/

参数解释-a ： 列出所有的文件与目录容量，因为默认仅统计目录的容量而已
-h: 以人们较易读的容量格式呈现(G/M/K)显示，自动选择显示的单位大小
-s : 列出总量而已，而不列出每个个别的目录占用容量
-k ： 以KB为单位进行显示
-m : 以MB为单位进行显示常用命令参考 查看当前目录大小[plain] du -sh ./
#######传输文件 end ############

#######关机 start ############
shutdown -h now  现在马上关机
shutdown -h 20:30  晚上8:30定时关机
shutdown -r now  现在马上重起
shutdown -r 20:30  晚上8:30定时重起
#######关机 end ############
```


