package handlers

import (
	handlers "github.com/danmaina/HttpResponse"
	"net/http"
)

func returnResponse(status int, err error, body interface{}, res http.ResponseWriter) {
	_ = handlers.Response{
		Status: status,
		Error:  err,
		Body:   body,
	}.ReturnResponse(res)
}
