package response

import (
	"encoding/json"
	"net/http"

	"github.com/skydrive/beans"
	"github.com/skydrive/config"
)

func ReturnResponseCodeMessage(w http.ResponseWriter, code int32, message string) {
	response := NewResponse(code, message)
	sonResult, _ := json.Marshal(response)
	setResult(w, http.StatusOK, sonResult)
}

func setResult(w http.ResponseWriter, code int, result []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, token")
	w.WriteHeader(code)
	w.Write(result)
}

func ReturnResponseCodeMessageHttpCode(w http.ResponseWriter, httpCode int, code int32, message string) {
	response := NewResponse(code, message)
	sonResult, _ := json.Marshal(response)
	setResult(w, httpCode, sonResult)
}

func ReturnMetaInfo(w http.ResponseWriter, code int32, message string, filemeta *beans.File) {
	//response := NewMetaInfoBaseResponse(code, message, filemeta)
	jsonResult, error := json.Marshal(&FormatResponse{
		Code:    code,
		Message: message,
		Data:    filemeta,
	})
	if error != nil {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "internel server error ")
		return
	}
	setResult(w, http.StatusOK, jsonResult)
}

func ReturnResponse(w http.ResponseWriter, code int32, message string, data interface{}) {
	jsonResult, error := json.Marshal(&FormatResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
	if error != nil {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "internel server error ")
		return
	}
	setResult(w, http.StatusOK, jsonResult)
}

func ReturnResponsePage(w http.ResponseWriter, code int32, message string, data interface{}, pageNo, pageSize, nextPageId, total int64) {
	jsonResult, error := json.Marshal(&FormatResponse{
		Code:    code,
		Message: message,
		Data: &PageData{
			NextPageId: nextPageId,
			PageNo:     pageNo,
			PageSize:   pageSize,
			Total:      total,
			Data:       data,
		},
	})
	if error != nil {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "internel server error ")
		return
	}
	setResult(w, http.StatusOK, jsonResult)
}

//构造
func NewResponse(code int32, message string) *FormatResponse {
	return &FormatResponse{
		Code:    code,
		Message: message,
	}
}
