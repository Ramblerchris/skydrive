package response
/*
type BaseResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}*/

type FormatResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PageData struct {
	NextPageId int64       `json:"nextpageid,omitempty"`
	PageNo     int64       `json:"pageNo"`
	PageSize   int64       `json:"pageSize"`
	Total      int64       `json:"total"`
	Data       interface{} `json:"list"`
}

