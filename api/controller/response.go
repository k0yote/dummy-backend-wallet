package controller

import "time"

type ApiResponse struct {
	ResponseTime  int64       `json:"responseTime"`
	StatusMessage string      `json:"statusMessage"`
	ResponseData  interface{} `json:"responseData,omitempty"`
}

const EMPRTY_MSG = ""

var (
	success = "Success"
)

func NewAPIResponse(statusMessage string, responseData interface{}) *ApiResponse {

	if len(statusMessage) == 0 {
		statusMessage = success
	}

	return &ApiResponse{
		ResponseTime:  time.Now().Local().Unix(),
		StatusMessage: statusMessage,
		ResponseData:  responseData,
	}
}
