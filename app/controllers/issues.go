/* ページ表示用コントローラー 処理はすべてAPIで行なう  */

package controllers

import (
	"fmt"
	"security-cop/app/models"
	"strconv"

	"github.com/revel/revel"
)

type Issues struct {
	*revel.Controller
}

func (c Issues) Show(id int) revel.Result {
	fmt.Printf("%d\n", id)
	rows, _ := Dbm.Select(models.Issue{}, "select * from issue where id="+strconv.Itoa(id))
	for _, row := range rows {
		issue := row.(*models.Issue)
		fmt.Printf("%d, %s\n", issue.Id, issue.Title)
		return c.Render(issue)
	}
	return nil
}

func (c Issues) List() revel.Result {
	rows, _ := Dbm.Select(models.Issue{}, "select * from issue")
	var issues []*models.Issue
	for _, row := range rows {
		issues = append(issues, row.(*models.Issue))
	}
	return c.Render(issues)
}
