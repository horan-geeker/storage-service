package response

import (
    "encoding/json"
    "log"
)

const (
    SUCCESS       = 0
    ERROR_PARAM   = 1
    NOT_FIND_DATA = 3
)

type responseStruct struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

var errorMessageMap = map[int]string{
    SUCCESS:       "请求成功",
    ERROR_PARAM:   "参数错误",
    NOT_FIND_DATA: "未找到数据",
}

func Json(status int, message string, data interface{}) string {
    if message == "" {
        if val, ok := errorMessageMap[status]; ok {
            message = val
        }
    }
    response := responseStruct{
        Status:  status,
        Message: message,
        Data:    data,
    }
    responseStr, err := json.Marshal(response)
    if err != nil {
        log.Println(err.Error())
    }
    return string(responseStr)
}

func JsonParamError(message string) string {
    if message == "" {
        if val, ok := errorMessageMap[ERROR_PARAM]; ok {
            message = val
        }
    }
    response := responseStruct{
        Status:  ERROR_PARAM,
        Message: message,
        Data:    nil,
    }
    responseStr, err := json.Marshal(response)
    if err != nil {
        log.Println(err.Error())
    }
    return string(responseStr)
}

func JsonDone() string {
    response := responseStruct{
        Status:  SUCCESS,
        Message: "success",
        Data:    nil,
    }
    responseStr, err := json.Marshal(response)
    if err != nil {
        log.Println(err.Error())
    }
    return string(responseStr)
}
