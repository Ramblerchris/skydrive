package config

import (
	"encoding/json"
	"fmt"
	"github.com/skydrive/logger"
	"github.com/skydrive/utils"
	"os"
	"time"
)

var (
	Debug         = true
	DiskFileRoot  = "updisk"
	AlbumFileRoot = "upalbum"
	ThumbnailRoot = "thumbnail"
	LogDir        = "log"

)
//全局配置项
type ConfigInt struct {
	USER_NAME string `json:"mysqlDbUsername"`
	PASS_WORD string `json:"mysqlDbPassword"`
	HOST string `json:"mysqlDbHost"`
	PORT string `json:"mysqlDbPort"`
	DATABASE string `json:"mysqlDbName"`
	CHARSET string `json:"mysqlDbcharset"`
	MySqlSprintf string `json:"mysqlDbconSprintf"`
	DiskFileRoot string `json:"diskFileRoot"`
	AlbumFileRoot string `json:"albumFileRoot"`
	ThumbnailRoot string `json:"thumbnailRoot"`

	Http_ServeLocation int `json:"HttpPORT"`
	UDP_SERVER_ListenPORT  int `json:"UDP_SERVER_ListenPORT"`
	UDP_GroupSERVER_ListenPORT int `json:"UDP_GroupSERVER_ListenPORT"`
	UDP_GroupSERVER_Sendport int `json:"UDP_GroupSERVER_SendPORT"`
	UDP_SERVER_Sendport  int `json:"UDP_SERVER_SendPORT"`
	AdminManagerDir  string `json:"adminManagerDir"`
	LogDir  string `json:"LogDir"`
}

func (dbConfig *ConfigInt) GetDataSourceName() string {
	return fmt.Sprintf(dbConfig.MySqlSprintf, dbConfig.USER_NAME, dbConfig.PASS_WORD, dbConfig.HOST, dbConfig.PORT, dbConfig.DATABASE, 1000, 500, 500, dbConfig.CHARSET)
}

func (dbConfig *ConfigInt) String() string {
	return fmt.Sprintf("USER_NAME:%s PASS_WORD:%s  HOST: %s  PORT: %s  DATABASE: %s   CHARSET: %s  MySqlSprintf: %s ",
		dbConfig.USER_NAME, dbConfig.PASS_WORD, dbConfig.HOST, dbConfig.PORT, dbConfig.DATABASE, dbConfig.CHARSET, dbConfig.MySqlSprintf)
}

//当config_ini.json 不存在的时候，使用默认配置
const (
	USER_NAME    = "root"
	PASS_WORD    = "nihao@123456"
	HOST         = "localhost"
	PORT         = "3306"
	DATABASE     = "skydrive"
	CHARSET      = "utf8"
	//MySqlConfit  = "%s:%s@tcp(%s:%s)/%s?charset=%s"
	MySqlSprintf = "%s:%s@tcp(%s:%s)/%s?timeout=%dms&readTimeout=%dms&writeTimeout=%dms&charset=%s"
	//("%s:%s@tcp(%s:%d)/%s?timeout=%dms&readTimeout=%dms&writeTimeout=%dms&charset=utf8", "用户名", "密码", "hostip", 端口, "库名", 1000, 500, 500)//后面三个分别为连接超时，读超时，写超时
	Http_ServeLocation         = 9996
	UDP_SERVER_ListenPORT      = 8997
	UDP_GroupSERVER_ListenPORT = 8998
	UDP_GroupSERVER_Sendport = 8999
	UDP_SERVER_Sendport      = 8996

	CHUNK_Size = 7 * 1025 * 1025 //分块上传的大小 7Mb

	Token_ExpriseTime = time.Hour * 24 * 7 //7天过期

	Net_ErrorCode_Token_exprise = -99

	Net_ErrorCode = -1

	Net_SuccessCode = 100
	//重复操作成功
	Net_SuccessAginCode = 101
	Net_EmptyCode       = 1

	Regex_MobilePhone = `^(1[3|5|6|7|8|9][0-9]\d{4,8})$`
	Salt_MD5          = "&%)&%A3t8C"

	Thumbnail_Quality      = 30
	Thumbnail_index      = 3
	//用于压缩gif
	Thumbnail_fuzz_gif    = 5
	//宽高百分比
	Thumbnail_widthf      = 30
	LOG_FILE_NAME          = "logfile.log"
	configInt              = "config_ini.json"
	AdminManagerDir        = "html"
)

func Setup()  (dbconfig *ConfigInt  ){
	exists, _ := utils.PathExists(configInt)
	if exists {
		file, _ := os.Open("config_ini.json")
		defer file.Close()
		decoder := json.NewDecoder(file)
		decoder.Decode(&dbconfig)
		logger.Info("读取配置文件",dbconfig.String())
		if len(dbconfig.DiskFileRoot)>0{
			DiskFileRoot = dbconfig.DiskFileRoot
		}
		if len(dbconfig.AlbumFileRoot)>0{
			AlbumFileRoot = dbconfig.AlbumFileRoot
		}
		if len(dbconfig.LogDir)>0{
			LogDir = dbconfig.LogDir
		}
		if len(dbconfig.ThumbnailRoot)>0{
			ThumbnailRoot = dbconfig.ThumbnailRoot
		}
	} else {
		dbconfig = &ConfigInt{
			USER_NAME:    USER_NAME,
			PASS_WORD:    PASS_WORD,
			HOST:         HOST,
			PORT:         PORT,
			DATABASE:     DATABASE,
			CHARSET:      CHARSET,
			MySqlSprintf: MySqlSprintf,

			Http_ServeLocation:         Http_ServeLocation,
			UDP_SERVER_ListenPORT:      UDP_SERVER_ListenPORT,
			UDP_GroupSERVER_ListenPORT: UDP_GroupSERVER_ListenPORT,
			UDP_GroupSERVER_Sendport:   UDP_GroupSERVER_Sendport,
			UDP_SERVER_Sendport:        UDP_SERVER_Sendport,
			AdminManagerDir:        AdminManagerDir,
		}
		logger.Info("默认配置",dbconfig.String())
	}
	return dbconfig
}
