package response

import (
	"encoding/json"
	"github.com/jihanlugas/pos/constant"
	"github.com/labstack/echo/v4"
)

// SuccessResponse type for Success Response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Payload interface{} `json:"payload" swaggertype:"object"`
}

type Payload map[string]interface{}

func (e *Response) Error() string {
	return e.Message
}

func Success(code int, msg string, payload interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Status:  true,
		Payload: payload,
	}
}

func Error(code int, msg string, payload interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Status:  false,
		Payload: payload,
	}
}

func ErrorForce(code int, msg string, payload Payload) *Response {
	payload["forceLogout"] = true
	return &Response{
		Code:    code,
		Message: msg,
		Status:  false,
		Payload: payload,
	}
}

func (r *Response) SendJSON(c echo.Context) error {
	if js, err := json.Marshal(r); err != nil {
		panic(err)
	} else {
		c.Set(constant.Response, js)
		return c.Blob(r.Code, echo.MIMEApplicationJSONCharsetUTF8, js)
	}
}

func ValidationError(err error) *Payload {
	return &Payload{
		"listError": getListError(err),
	}
}
