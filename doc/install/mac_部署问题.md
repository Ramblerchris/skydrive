###### Mysql 安装

```shell
brew install mysql@5.7
```

设置path
按照上面运行的结果提示 ：

```shell

If you need to have openssl@1.1 first in your PATH, run:
  echo 'export PATH="/opt/homebrew/opt/openssl@1.1/bin:$PATH"' >> ~/.zshrc

For compilers to find openssl@1.1 you may need to set:
  export LDFLAGS="-L/opt/homebrew/opt/openssl@1.1/lib"
  export CPPFLAGS="-I/opt/homebrew/opt/openssl@1.1/include"

==> mysql@5.7
We've installed your MySQL database without a root password. To secure it run:
    mysql_secure_installation

MySQL is configured to only allow connections from localhost by default

To connect run:
    mysql -uroot

mysql@5.7 is keg-only, which means it was not symlinked into /opt/homebrew,
because this is an alternate version of another formula.

If you need to have mysql@5.7 first in your PATH, run:
  echo 'export PATH="/opt/homebrew/opt/mysql@5.7/bin:$PATH"' >> ~/.zshrc

For compilers to find mysql@5.7 you may need to set:
  export LDFLAGS="-L/opt/homebrew/opt/mysql@5.7/lib"
  export CPPFLAGS="-I/opt/homebrew/opt/mysql@5.7/include"


To have launchd start mysql@5.7 now and restart at login:
  brew services start mysql@5.7
Or, if you don't want/need a background service you can just run:
  /opt/homebrew/opt/mysql@5.7/bin/mysql.server start
```

先开启服务：

```shell
/opt/homebrew/opt/mysql@5.7/bin/mysql.server start
/opt/homebrew/opt/mysql@5.7/bin/mysql.server restart/start/stop
```

然后设置密码按照提示配置：

```shell
mysql_secure_installation
```
