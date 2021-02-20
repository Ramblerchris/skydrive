package handler

import (
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/handler/cache"
	"github.com/skydrive/logger"
	"github.com/skydrive/response"
	"github.com/skydrive/utils"
	"net/http"
	"time"
)

const TAG = "auth.go"

var tokenmap = cache.NewTokenMap()

func init() {
	//生产项目推荐用缓存策略，有值淘汰，或者直接使用redis

}

func BuildEncodePwd(pwd string) string {
	target := utils.GetStrMD5(pwd + config.Salt_MD5)
	logger.Infof("%s encode md5 %s\n", pwd, target)
	return target
}

type HandlerFuncAuth func(http.ResponseWriter, *http.Request, *db.TableUToken)

//测试网络是否连通，
func CheckNetIsOkHandler(w http.ResponseWriter, r *http.Request) {
	response.ReturnResponseCodeMessage(w, config.Net_SuccessCode, "连接成功")
	logger.Info(" CheckNetIsOkHandler :", " Now:", time.Now().UnixNano()/1e6)
}

func getToken(r *http.Request) string {
	token := r.FormValue("token")
	if len(token) == 0 || token == "" {
		token = r.Header.Get("token")
	}
	return token
}

func TokenCheckInterceptor(h HandlerFuncAuth) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("request URL", r.Method, r.RequestURI)

		if r.Method == "OPTIONS" {
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Add("Access-Control-Allow-Origin", "*")
				w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, HEAD, OPTIONS")
				w.Header().Add("Access-Control-Allow-Headers", "token, Content-Type")
				w.Header().Set("Content-Type", "application/json;charset=utf-8")
			}
			return
		}
		r.ParseForm()
		//logger.Error("Content-Type",contenttype)
		if r.Method == "POST" {
			r.ParseMultipartForm(32 << 20)
		}
		token := getToken(r)
		/*token := r.FormValue("token")
		if len(token) == 0 || token == "" {
			if r.Method == "POST" {
				token = r.PostFormValue("token")
				if len(token) == 0 || token == "" {
					r.ParseMultipartForm(32 << 20)
					if r.MultipartForm != nil {
						values := r.MultipartForm.Value["token"]
						if len(values) > 0 {
							token= values[0]
						}
					}
				}

			}
		}*/
		if len(token) == 0 || token == "" {
			response.ReturnResponseCodeMessageHttpCode(w, http.StatusUnauthorized, config.Net_ErrorCode, "bad request")
			return
		}
		userTokenInfo, exist := tokenmap.ReadTokenMap(token)
		if !exist {
			if byTokenbydb, er := db.GetUserTokenInfoByToken(token); er != nil {
				response.ReturnResponseCodeMessageHttpCode(w, http.StatusForbidden, config.Net_ErrorCode, "bad request")
				return
			} else {
				userTokenInfo = byTokenbydb
				tokenmap.WriteTokenMap(token, userTokenInfo)
			}
		} else {
			logger.Info(TAG, "缓存获取用户:", userTokenInfo)
		}
		if userTokenInfo.User_token.String != "" {
			//todo 判断过期时间
			logger.Info(token, " Expiretime :", userTokenInfo.Expiretime.Int64, " Now:", time.Now().UnixNano()/1e6)
			if userTokenInfo.Expiretime.Int64 < time.Now().UnixNano()/1e6 {
				tokenmap.DeleteTokenMap(token)
				response.ReturnResponseCodeMessageHttpCode(w, http.StatusForbidden, config.Net_ErrorCode_Token_exprise, "token expired")
				return
			}
			h(w, r, &userTokenInfo)
			return
		} else {
			response.ReturnResponseCodeMessageHttpCode(w, http.StatusForbidden, config.Net_ErrorCode, "bad request")
		}
	}
}
