package response

import (
	"encoding/json"
	"github.com/skydrive/config"
	"github.com/skydrive/handler"
	"net/http"
)

func ReturnResponseCodeMessage(w http.ResponseWriter, code int32, message string) {
	w.Header().Add("Content-Type", "application/json")
	response := NewResponse(code, message)
	sonResult, _ := json.Marshal(response)
	// w.WriteHeader(http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
	w.Write(sonResult)
}

func ReturnResponseCodeMessageHttpCode(w http.ResponseWriter,httpCode int, code int32, message string) {
	w.Header().Add("Content-Type", "application/json")
	response := NewResponse(code, message)
	sonResult, _ := json.Marshal(response)
	// w.WriteHeader(http.StatusInternalServerError)
	w.WriteHeader(httpCode)
	w.Write(sonResult)
}


func ReturnMetaInfo(w http.ResponseWriter, code int32, message string, filemeta *handler.UserFile) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := NewMetaInfoBaseResponse(code, message, filemeta)
	jsonResult, error := json.Marshal(response)
	if error != nil {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "internel server error ")
		return
	}
	w.Write(jsonResult)
}

func ReturnResponse(w http.ResponseWriter, code int32, message string, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResult, error := json.Marshal(&FormatResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
	if error != nil {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "internel server error ")
		return
	}
	w.Write(jsonResult)
}

func ReturnResponsePage(w http.ResponseWriter, code int32, message string, data interface{},pageNo ,pageSize ,nextPageId,total int64) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResult, error := json.Marshal(&FormatResponse{
		Code:    code,
		Message: message,
		Data: &PageData{
			NextPageId:   nextPageId,
			PageNo:   pageNo,
			PageSize: pageSize,
			Total:    total,
			Data:     data,
		},
	})
	if error != nil {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "internel server error ")
		return
	}
	w.Write(jsonResult)
}

