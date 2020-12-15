package response

import "github.com/skydrive/meta"

type BaseResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

//组合方式
type MetaInfoBaseResponse struct {
	BaseResponse
	Data meta.FileMeta `json:"data"`
}

type FormatResponse struct {
	Code    int32  	`json:"code"`
	Message string 	`json:"message"`
	Data 	interface{}	`json:"data"`
}

type PageData struct {
	PageNo    int64  `json:"pageNo"`
	PageSize  int64  `json:"pageSize"`
	Total	  int64  `json:"total"`
	Data  	  interface{} `json:"list"`
}

//构造
func NewResponse(code int32, message string) *BaseResponse {
	return &BaseResponse{
		code,
		message,
	}
}

//构造
func NewMetaInfoBaseResponse(code int32, message string, meta *meta.FileMeta) *MetaInfoBaseResponse {
	metaInfoBaseResponse := &MetaInfoBaseResponse{}
	metaInfoBaseResponse.Code = code
	// metaInfoBaseResponse.BaseResponseData.Message=message
	metaInfoBaseResponse.Message = message
	if meta != nil {
		metaInfoBaseResponse.Data = *meta
	}
	return metaInfoBaseResponse
}

