package response

import (
	"encoding/json"
	"fmt"
	"log"
)

func (response *FormatResponse) GetResponseCodeMessage(code int32, message string) {
	response.Code = code
	response.Message = message
}

func (response *FormatResponse) GetResponse(code int32, message string,data interface{}) {
	response.Code = code
	response.Message = message
	response.Data = data
}

func (response *BaseResponse) GetResponseBytes()[]byte {
	r, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	return r
}

func (response *BaseResponse) GetResponseStr()string {
	r, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	return string(r)
}

func GetJsonBytesCodeMessage(code int ,msg string ) []byte {
	return []byte(fmt.Sprintf(`{"code":%d,"msg":"%s"}`,code ,msg))
}

func GetJsonStrCodeMessage(code int ,msg string)string  {
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`,code ,msg)
}

