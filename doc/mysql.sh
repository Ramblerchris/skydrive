#login mysql
mysql -h192.168.1.105 -uroot -p
create database skydrive;
show databases;
use skydrive
drop table tbl_file
drop table tbl_user_file
drop table tbl_user
#删除表
drop table tbl_user_token
#查看表结构
desc tbl_user_token


#
#1.Windows下

#启动服务
mysqld --console
或 net start mysql
#关闭服务
mysqladmin -uroot shudown
或 net stop mysql


#2.Linux下

#启动服务
service mysql start
#关闭服务
service mysql stop
#重启服务
service restart stop