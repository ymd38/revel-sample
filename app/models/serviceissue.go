package models

import (
	. "security-cop/app/util"
	"strconv"
	"sync"
	"time"
)

type ServiceIssue struct {
	GorpController
}

func (si *ServiceIssue) Create(issueid int, serviceid int) error {
	si_data := &ServiceIssueData{ID: 0, IssueID: issueid, ServiceID: serviceid}
	si_data.Status = STATUS_NOCHECK
	si_data.Created = time.Now().Unix()
	si_data.Updated = time.Now().Unix()

	if err := si_data.Validate(); err != nil {
		return err
	}

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
		}(issueid, service_list[i].ID)
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
		}(issue_list[i].ID, serviceid)
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
		si_list[cnt].ID = sidata.ID
		si_list[cnt].ServiceID = sidata.ServiceID
		si_list[cnt].IssueID = sidata.IssueID
		si_list[cnt].Status = sidata.Status
		si_list[cnt].Memo = sidata.Memo
		si_list[cnt].ReflectdateStr = UnixTimeToDayString(sidata.Reflectdate)
		si_list[cnt].CreatedStr = UnixTimeToDateString(sidata.Created)
		si_list[cnt].UpdatedStr = UnixTimeToDateString(sidata.Updated)
		cnt++
	}
	return si_list
}
