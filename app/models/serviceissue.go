package models

import (
	"errors"
	"modeltest/app/util"
	"time"
)

type ServiceIssue struct {
	GorpController
}

func (si *ServiceIssue) Create(si_data *ServiceIssueData) error {
	/*	var v revel.Validation
		service_data.Validate(&v)
		if v.HasErrors() {
			return errors.New("Validate Error")
		}
	*/
	si_data.Status = util.STATUS_NOCHECK
	//gorp doesn't support time type. we use unix time on DB.
	si_data.Created = time.Now().Unix()
	si_data.Updated = time.Now().Unix()

	err := Txn.Insert(si_data)
	if err != nil {
		panic(err)
		return errors.New("System Error")
	}
	return nil
}
