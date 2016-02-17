package models

import (
	"fmt"
	"modeltest/app/util"
	"sync"
	"time"
)

type ServiceIssue struct {
	GorpController
}

func (si *ServiceIssue) Create(issueid int, service_list []ServiceData) error {
	var wg sync.WaitGroup
	for i := 0; i < len(service_list); i++ {
		fmt.Println(service_list[i])
		wg.Add(1)
		go func(iid int, sid int) {
			fmt.Println(iid, sid)
			defer wg.Done()
			si_data := &ServiceIssueData{Id: 0, IssueId: iid, ServiceId: sid}
			si_data.Status = util.STATUS_NOCHECK
			si_data.Created = time.Now().Unix()
			si_data.Updated = time.Now().Unix()
			err := Txn.Insert(si_data)
			if err != nil {
				panic(err)
			}
		}(issueid, service_list[i].Id)
	}
	wg.Wait()
	return nil
}
