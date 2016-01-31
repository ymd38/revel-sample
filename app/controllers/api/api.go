/* API共通機能 */
package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"security-cop/app/controllers"

	"github.com/revel/revel"
)

type ApiController struct {
	controllers.App
	controllers.GorpController
}

type Response struct {
	Code    int         `json:"code"`
	Results interface{} `json:"results,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	OK int = iota
	WARN_NOT_FOUND
	ERR_VALIDATE
	ERR_SYSTEM
)

func init() {
	//authはinit.goで行うので不要かも
	//revel.InterceptMethod(controllers.App.Auth, revel.BEFORE)
	revel.InterceptMethod((*controllers.GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*controllers.GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*controllers.GorpController).Rollback, revel.FINALLY)
}

func (c *ApiController) BindParams(s interface{}) error {
	return JsonDecode(c.App.Request.Body, s)
}

func JsonDecode(i io.Reader, s interface{}) error {
	bytes, err := ioutil.ReadAll(i)
	if err != nil {
		return nil
	}

	if len(bytes) == 0 {
		return nil
	}

	return json.Unmarshal(bytes, s)
}

func (c *ErrorResponse) String() string {
	return ErrorMessage(c.Code)
}

func ErrorMessage(code int) string {
	switch code {
	case WARN_NOT_FOUND:
		return "not found"
	case ERR_VALIDATE:
		return "validate error"
	case ERR_SYSTEM:
		return "system error"
	default:
		return "system error"
	}
}
