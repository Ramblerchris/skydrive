module github.com/skydrive

go 1.13

replace github.com/skydrive => ../skydrive

require (
	//配置文件的使用由来已久，从.ini、XML、JSON、YAML再到TOML，语言的表达能力越来越强，同时书写便捷性也在不断提升。
	//TOML是前GitHub CEO， Tom Preston-Werner，于2013年创建的语言，其目标是成为一个小规模的易于使用的语义化配置文件格式。
	//TOML被设计为可以无二义性的转换为一个哈希表(Hash table)。
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	//github.com/esimov/caire v1.2.6 // indirect
	//github.com/esimov/pigo v1.4.2 // indirect
	//github.com/h2non/bimg v1.1.4 // indirect
	github.com/go-redis/redis/v8 v8.3.3
	github.com/go-sql-driver/mysql v1.5.0
	//图片压缩工具
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	//Package pretty提供Go值的漂亮打印。
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	//UUid
	github.com/satori/go.uuid v1.2.0
	//获取系统信息
	//github.com/shirou/gopsutil v3.20.12+incompatible
	github.com/shirou/gopsutil v2.20.8+incompatible
	//日志
	github.com/sirupsen/logrus v1.7.0
	//文件解析及手动实现ini文件解析
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
