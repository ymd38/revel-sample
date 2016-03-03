package models

import (
	"errors"
	. "security-cop/app/util"
	"strconv"
	"time"

	"github.com/revel/revel"
)

type Service struct {
	GorpController
}

func (service *Service) Create(service_data *ServiceData) error {
	var v revel.Validation
	if err := service_data.Validate(&v); err != nil {
		return err
	}

	service_data.Start = DayStringToUnixTime(service_data.StartStr)
	service_data.End = DayStringToUnixTime(service_data.EndStr)
	//gorp doesn't support time type. we use unix time on DB.
	service_data.Created = time.Now().Unix()
	service_data.Updated = time.Now().Unix()

	if err := Txn.Insert(service_data); err != nil {
		return errors.New("Insert Error")
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
	service_list := make([]ServiceData, len(rows))
	cnt := 0
	for _, row := range rows {
		servicedata := row.(*ServiceData)
		service_list[cnt].Id = servicedata.Id
		service_list[cnt].Name = servicedata.Name
		if servicedata.Start != 0 {
			service_list[cnt].StartStr = UnixTimeToDayString(servicedata.Start)
		}

		if servicedata.End != 0 {
			service_list[cnt].EndStr = UnixTimeToDayString(servicedata.End)
		}

		service_list[cnt].CreatedStr = UnixTimeToDateString(servicedata.Created)
		service_list[cnt].UpdatedStr = UnixTimeToDateString(servicedata.Updated)
		cnt++
	}
	return service_list
}
