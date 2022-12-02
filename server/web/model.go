package web

import "errors"

// Error define.
var ErrRunEngineFail = errors.New("ERR_RUN_ENGINE_FAIL")
var ErrBindBodyFail = errors.New("ERR_BIND_BODY_FAIL")
var ErrAuthFail = errors.New("ERR_AUTH_FAIL")
var ErrInternal = errors.New("ERR_INTERNAL")

// Request model.
// Create new notes request.
type CreateNewNoteReq struct {
	AccessKey string `json:"access_key"`
	Category  string `json:"category"`
	Data      string `json:"data"`
}

// List notes request.
type ListNoteReq struct {
	AccessKey string `json:"access_key"`
	Limit     uint   `jsone:"limit"` // The most number of the query notes.
}

// List notes by category request.
type ListNoteByCategoryReq struct {
	AccessKey string `json:"access_key"`
	Category  string `json:"category"` // The category of the query notes.
	Limit     uint   `jsone:"limit"`   // The most number of the query notes.
}

// Response model.
// The error response model.
type ErrResp struct {
	ErrMsg string `json:"error_msg"`
}

// The success response model.
type SuccessResp struct {
	ErrMsg string `json:"error_msg"`
	Data   interface{}
}
