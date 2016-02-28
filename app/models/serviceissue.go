package models

import (
	"security-cop/app/util"
	"strconv"
	"sync"
	"time"
)

type ServiceIssue struct {
	GorpController
}

func (si *ServiceIssue) Create(issueid int, serviceid int) error {
	si_data := &ServiceIssueData{Id: 0, IssueId: issueid, ServiceId: serviceid}
	si_data.Status = util.STATUS_NOCHECK
	si_data.Created = time.Now().Unix()
	si_data.Updated = time.Now().Unix()
	err := Txn.Insert(si_data)
	if err != nil {
		panic(err)
	}
	return nil
}

func (si *ServiceIssue) CreateByIssueID(issueid int) error {
	var wg sync.WaitGroup
	service := new(Service)
	service_list := service.GetList()

	for i := 0; i < len(service_list); i++ {
		wg.Add(1)
		go func(issueid int, serviceid int) {
			defer wg.Done()
			data := si.getList("where issueid=" + strconv.Itoa(issueid) + " and serviceid=" + strconv.Itoa(serviceid))
			if len(data) == 0 {
				si.Create(issueid, serviceid)
			}
		}(issueid, service_list[i].Id)
	}
	wg.Wait()
	return nil
}

func (si *ServiceIssue) CreateByServiceID(serviceid int) error {
	var wg sync.WaitGroup
	issue := new(Issue)
	issue_list := issue.GetList() //TODO:抽出条件を設定

	for i := 0; i < len(issue_list); i++ {
		wg.Add(1)
		go func(issueid int, serviceid int) {
			defer wg.Done()
			data := si.getList("where issueid=" + strconv.Itoa(issueid) + " and serviceid=" + strconv.Itoa(serviceid))
			if len(data) == 0 {
				si.Create(issueid, serviceid)
			}
		}(issue_list[i].Id, serviceid)
	}
	wg.Wait()
	return nil
}

//TODO:gorpからgormに切り替える
func (si *ServiceIssue) getList(condition string) []ServiceIssueData {
	sql := "select * from service_issue " + condition
	rows, _ := Dbm.Select(ServiceIssueData{}, sql)
	si_list := make([]ServiceIssueData, len(rows))
	cnt := 0
	for _, row := range rows {
		sidata := row.(*ServiceIssueData)
		si_list[cnt].Id = sidata.Id
		si_list[cnt].ServiceId = sidata.ServiceId
		si_list[cnt].IssueId = sidata.IssueId
		si_list[cnt].Status = sidata.Status
		si_list[cnt].Memo = sidata.Memo
		si_list[cnt].ReflectdateStr = util.UnixTimeToDayString(sidata.Reflectdate)
		si_list[cnt].CreatedStr = util.UnixTimeToDateString(sidata.Created)
		si_list[cnt].UpdatedStr = util.UnixTimeToDateString(sidata.Updated)
		cnt++
	}
	return si_list
}
