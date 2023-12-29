package handler

import "github.com/TravisRoad/goshower/internal/model"

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type AuthResponse struct {
	BaseResponse
	Data model.UserInfo `json:"data"`
}
