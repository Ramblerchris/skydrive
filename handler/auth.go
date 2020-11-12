package handler

import (
	"fmt"
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/utils"
	"net/http"
	"time"
)

func BuildEncodePwd(pwd string) string {
	target := utils.GetStrMD5(pwd + config.Salt_MD5)
	fmt.Printf("%s encode md5 %s\n", pwd, target)
	return target
}

type HandlerFuncAuth func(http.ResponseWriter, *http.Request, *db.UToken)

func CheckNet(w http.ResponseWriter, r *http.Request) {
	ReturnResponseCodeMessage(w, config.Net_SuccessCode, "连接成功")
	fmt.Println(" CheckNet :", " Now:", time.Now().UnixNano()/1e6)
}

func getToken(r *http.Request) string {
	token := r.FormValue("token")

	if len(token) == 0 || token == "" {
		token = r.Header.Get("token")
	}
	fmt.Println(" get token  :", token)

	return token
}

func TokenCheckInterceptor(h HandlerFuncAuth) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//fmt.Println("Content-Type",contenttype)
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
			//ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
			ReturnResponseCodeMessageHttpCode(w, http.StatusUnauthorized, config.Net_ErrorCode, "bad request")
			return
		}
		if byToken, err := db.GetUTokenByToken(token); err == nil && byToken.User_token.String != "" {
			//todo 判断过期时间
			fmt.Println(" Expiretime :", byToken.Expiretime.Int64, " Now:", time.Now().UnixNano()/1e6)
			if byToken.Expiretime.Int64 < time.Now().UnixNano()/1e6 {
				ReturnResponseCodeMessageHttpCode(w, http.StatusForbidden, config.Net_ErrorCode_Token_exprise, "token expired")
				return
			}
			h(w, r, &byToken)
			return
		}
		ReturnResponseCodeMessageHttpCode(w, http.StatusForbidden, config.Net_ErrorCode, "bad request")
	}
}
