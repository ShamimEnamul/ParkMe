package utils

import (
	"encoding/json"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/server/web/context"
)

type APIResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int      `json:"status"`
	Message []string `json:"message"`
}

var Response = make(map[string]APIResponse)

func init() {

	Response["catch_error"] = APIResponse{4000, "Error from catch block"}
	Response["success"] = APIResponse{1000, "success"}
	Response["failed"] = APIResponse{9999, "failed"}
	Response["SQL_Error"] = APIResponse{4001, "Syntax error"}
	Response["validation_error"] = APIResponse{2008, "Validation error"}
	Response["instance_creation_error"] = APIResponse{2008, "error occurred while creating new instance"}
}

func GetCode(message string) APIResponse {
	return Response[message]
}

func GetValidationCode(message []string) ErrorResponse {
	return ErrorResponse{Status: 2008, Message: message}
}

func ServeJSON(ctx *context.Context, err error, message APIResponse, logName string) bool {
	if err != nil {
		logs.Error(logName+": ", err)
		ctx.Output.SetStatus(200)
		resBody, _ := json.Marshal(message)
		ctx.Output.Body(resBody)
		return true
	}

	return false
}
