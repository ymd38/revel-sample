package models

import (
	. "security-cop/app/util"
	"strconv"
	"time"
)

type Service struct {
	GorpController
}

func (service *Service) Create(serviceData *ServiceData) error {
	if err := serviceData.Validate(); err != nil {
		return err
	}

	if serviceData.StartStr != "" {
		serviceData.Start = DayStringToUnixTime(serviceData.StartStr)
	}

	if serviceData.EndStr != "" {
		serviceData.End = DayStringToUnixTime(serviceData.EndStr)
	}

	//gorp doesn't support time type. we use unix time on DB.
	serviceData.Created = time.Now().Unix()
	serviceData.Updated = time.Now().Unix()

	if err := Txn.Insert(serviceData); err != nil {
		return err
	}

	return nil
}

func (service *Service) GetList() []ServiceData {
	return service.getList("")
}

func (service *Service) GetByID(id int) []ServiceData {
	return service.getList("where id=" + strconv.Itoa(id))
}

//TODO:今度gormにする
//privateで呼ばれるリスト取得
func (service *Service) getList(condition string) []ServiceData {
	sql := "select * from service " + condition
	rows, _ := Dbm.Select(ServiceData{}, sql)
	serviceList := make([]ServiceData, len(rows))
	cnt := 0
	for _, row := range rows {
		serviceData := row.(*ServiceData)
		serviceList[cnt].ID = serviceData.ID
		serviceList[cnt].Name = serviceData.Name
		if serviceData.Start != 0 {
			serviceList[cnt].StartStr = UnixTimeToDayString(serviceData.Start)
		}

		if serviceData.End != 0 {
			serviceList[cnt].EndStr = UnixTimeToDayString(serviceData.End)
		}

		serviceList[cnt].CreatedStr = UnixTimeToDateString(serviceData.Created)
		serviceList[cnt].UpdatedStr = UnixTimeToDateString(serviceData.Updated)
		cnt++
	}
	return serviceList
}
