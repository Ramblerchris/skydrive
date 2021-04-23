
go build 编译程序时可以通过 -ldflags 来指定编译参数。

-s 的作用是去掉符号信息。 -w 的作用是去掉调试信息。

测试加与不加 -ldflags 编译出的应用大小。

打包：

```shell
go build -ldflags "-s -w"
```




UPX 压缩
https://github.com/upx/upx/releases

windows:
https://github.com/upx/upx/releases

mac:

```shell
brew install upx
```

linux:

```shell
tar -Jxf upx*.tar.xz
```


压缩命令 upx -9 -o ./frpc2_upx ./frpc2

-o 指定压缩后的文件名。-9指定压缩级别，1-9。