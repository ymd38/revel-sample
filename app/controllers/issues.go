/* ページ表示用コントローラー 処理はすべてAPIで行なう  */

package controllers

import "github.com/revel/revel"

type Issues struct {
	*revel.Controller
}

func (c *Issues) Regist() revel.Result {
	return c.Render()
}
