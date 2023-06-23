package services

import (
	"goPandoraAdmin-Server/model"
)

// RespondHandle 接口返回处理函数
func RespondHandle(code int, msg interface{}, data interface{}) model.RespondStruct {
	var resp model.RespondStruct

	var message *string
	if msg != nil {
		switch v := msg.(type) {
		case string:
			message = &v
		case error:
			e := v.Error()
			message = &e
		}
	}
	switch code {
	case 0:
		resp.Status = "success"
		resp.Data = data
		resp.Message = message
	case -1:
		resp.Status = "error"
		resp.Message = message
	}
	return resp
}
