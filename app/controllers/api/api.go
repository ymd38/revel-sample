/* API共通機能 */
package controllers

import "security-cop/app/controllers"

const CALLBACK = "callBack"

type ApiController struct {
	controllers.App
}

type Response struct {
	Code    int         `json:"code"`
	Results interface{} `json:"results,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/* api response code */
const (
	OK int = iota
	WARN_NOT_FOUND
	ERR_VALIDATE
	ERR_FATAL
)
