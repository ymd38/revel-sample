package models

import (
	"errors"
	"security-cop/app/util"
	"time"

	"github.com/revel/revel"
)

type Service struct {
	GorpController
}

func (service *Service) Create(service_data *ServiceData) error {
	var v revel.Validation
	service_data.Validate(&v)
	if v.HasErrors() {
		return errors.New("Validate Error")
	}

	service_data.Start = util.DayStringToUnixTime(service_data.StartStr)
	service_data.End = util.DayStringToUnixTime(service_data.EndStr)
	//gorp doesn't support time type. we use unix time on DB.
	service_data.Created = time.Now().Unix()
	service_data.Updated = time.Now().Unix()

	err := Txn.Insert(service_data)
	if err != nil {
		return errors.New("System Error")
	}

	return nil
}

//common function for get services
func (service *Service) GetServiceList(condition string) []ServiceData {
	sql := "select * from service " + condition
	rows, _ := Dbm.Select(ServiceData{}, sql)
	service_list := make([]ServiceData, len(rows))
	cnt := 0
	for _, row := range rows {
		servicedata := row.(*ServiceData)
		service_list[cnt].Id = servicedata.Id
		service_list[cnt].Name = servicedata.Name
		service_list[cnt].StartStr = util.UnixTimeToDayString(servicedata.Start)
		service_list[cnt].EndStr = util.UnixTimeToDayString(servicedata.End)
		service_list[cnt].CreatedStr = util.UnixTimeToDateString(servicedata.Created)
		service_list[cnt].UpdatedStr = util.UnixTimeToDateString(servicedata.Updated)
		cnt++
	}
	return service_list
}
