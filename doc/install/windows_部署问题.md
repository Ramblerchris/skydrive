###### Mysql 安装

配置环境path变量：D:\Program Files\mysql-8.0.23-winx64\bin

在D:\Program Files\mysql-8.0.23-winx64 目录下新建 my.ini

```ini
[mysql]
# 设置mysql客户端默认字符集
default-character-set=utf8 
[mysqld]
#设置3306端口
port = 3306 
# 设置mysql的安装目录
basedir=D:\\Program Files\\mysql-8.0.23-winx64
# set datadir to the location of your data directory 设置mysql数据库的数据的存放目录
datadir=D:/Mysql/data
# 允许最大连接数
max_connections=200
# 服务端使用的字符集默认为8比特编码的latin1字符集
character-set-server=utf8
# 创建新表时将使用的默认存储引擎
default-storage-engine=INNODB
```

cmd 运行

```shell
mysqld --initialize --console
```

重新初始化一个data文件 D:/Mysql/data 文件夹会自动创建，不要手动创建

会看到临时密码

cmd 运行

```shell
  mysqld --install   出现 Service successfully installed.即为成功！
```

开启服务：

net start mysql



idea连接mysql报错Server returns invalid timezone. Go to 'Advanced' tab and set 'serverTimezone' property

```shell
# 设置全局时区 mysql> set global time_zone = '+8:00';
Query OK, 0 rows affected (0.00 sec) 
# 设置时区为东八区 mysql> set time_zone = '+8:00'; 
Query OK, 0 rows affected (0.00 sec) 
# 刷新权限使设置立即生效 mysql> flush privileges; 
Query OK, 0 rows affected (0.00 sec)
mysql> show variables like '%time_zone%';
```





问题：

- dLL缺少：
  https://download.visualstudio.microsoft.com/download/pr/366c0fb9-fe05-4b58-949a-5bc36e50e370/015EDD4E5D36E053B23A01ADB77A2B12444D3FB6ECCEFE23E3A8CD6388616A16/VC_redist.x64.exe

- 服务不能正常启动：
  data文件夹是自己创建的，删除data文件夹，在DOS界面进入到MySQL的文件夹下输入这个命令：mysqld  --initialize ，重新初始化一个data文件

- 生成临时密码：
  mysqld --initialize --console     第一次登陆MYSQL时，会提示要求输入初始密码，这是考虑安全因素，命令：mysqld –initialize会随机生成密码。初始密码在上图data文件夹下的xxx.err文件中，可以用记事本打开，用ctrl+f 查找功能找到如下一行记录：[Note] A temporary password is generated for root@localhost: NZ+uhXPq1zN.   其中NZ+uhXPq1zN.即为初始密码（注意.号不要漏了）进入后可以用如下命令修改，这里密码改为root：ALTER USER 'root'@'localhost' IDENTIFIED BY '密码';