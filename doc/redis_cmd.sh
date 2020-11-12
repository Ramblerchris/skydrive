//先启动server
redis-server

//连接
redis-cli
//指定ip 登录
redis-cli -h 127.0.0.1 -p 6379 -a 123456

//登录密码,如果设置过密码的情况下
 auth 123456

//查看所有keys
 keys *

//查看密码
config get requirepass
1) "requirepass"
2) ""

#报错： redis client init error ERR Client sent AUTH, but no password is set
#//需要设置密码
config set requirepass 123456



#window  启动直接点击 redis-server.exe

