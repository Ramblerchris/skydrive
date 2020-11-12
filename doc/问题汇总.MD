#### **问题**：

```sh
# command-line-arguments
./main.go:12:14: undefined: getIntranetIp
```

##### *==原因：==*

```sh
main 包中的不同的文件的代码不能相互调用，其他包可以。所以其实Link.go没有被一起编译执行。
```

##### *==解决:==*

```sh
如果是用命令行环境的，采用命令
go run *.go
```
#### **问题**：

```sh
 runtime error: invalid memory address or nil pointer dereference
```

##### *==原因：==*

```go
//当我们在使用 := 时 会创建一个新的mysqldb变量,新的sqlDb会把全局变量sqlDb覆盖掉
var mysqldb *sql.DB
func init() {
	mysqldb, error := sql.Open("mysql", config.GetDataSourceName())
}
```

##### *==解决:==*

```go
var mysqldb *sql.DB
func init() {
	var error error
	mysqldb, error = sql.Open("mysql", config.GetDataSourceName())
}
```

##### 

#### **问题**：

```sh
api url中多一个/ method 会变成get 
例如
http://localhost:9996//userfile/uploadmultipartinfo
```

##### *==解决:==*

```go
去掉多余的斜杠
```

##### 

