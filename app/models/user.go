package models

import (
	"fmt"
	. "security-cop/app/util"
	"time"
)

type User struct {
	GorpController
}

func (user *User) Create(userData *UserData) error {
	if err := userData.Validate(); err != nil {
		return err
	}

	userData.Password = ToMD5(userData.Password)
	fmt.Println(userData.Password)
	//gorp doesn't support time type. we use unix time on DB.
	userData.Created = time.Now().Unix()
	userData.Updated = time.Now().Unix()

	if err := Txn.Insert(userData); err != nil {
		fmt.Println(err)
		return err
	}

	userData.Password = ""

	return nil
}

/*
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
*/
