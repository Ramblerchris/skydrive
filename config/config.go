package config

import (
	"fmt"
	"time"
)

const (
	USER_NAME                  = "root"
	PASS_WORD                  = "nihao@123456"
	HOST                       = "localhost"
	PORT                       = "3306"
	DATABASE                   = "skydrive"
	CHARSET                    = "utf8"
	MySqlConfit                = "%s:%s@tcp(%s:%s)/%s?charset=%s"
	MySqlConfit1                = "%s:%s@tcp(%s:%s)/%s?timeout=%dms&readTimeout=%dms&writeTimeout=%dms&charset=%s"
//("%s:%s@tcp(%s:%d)/%s?timeout=%dms&readTimeout=%dms&writeTimeout=%dms&charset=utf8", "用户名", "密码", "hostip", 端口, "库名", 1000, 500, 500)//后面三个分别为连接超时，读超时，写超时
	ServeLocation              = ":9996"
	UDP_SERVER_ListenPORT      =8997
	UDP_GroupSERVER_ListenPORT =8998


	UDP_GroupSERVER_SendPORT =8999
	UDP_SERVER_SendPORT =8996


	CHUNK_Size= 7 * 1025 * 1025//分块上传的大小 7Mb


	Token_ExpriseTime=time.Hour*24*7  //7天过期


	Net_ErrorCode_Token_exprise   = -99

	Net_ErrorCode   = -1

	Net_SuccessCode = 100
	//重复操作成功
	Net_SuccessAginCode = 101
	Net_EmptyCode = 1

	Regex_MobilePhone=`^(1[3|4|5|8][0-9]\d{4,8})$`
	Salt_MD5="&%)&%A3t8C"

)

func GetDataSourceName() string {
	return fmt.Sprintf(MySqlConfit1, USER_NAME, PASS_WORD, HOST, PORT, DATABASE, 1000, 500, 500,CHARSET)
}
