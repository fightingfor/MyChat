package service

import (
	"MyChat/model"
	"io/ioutil"
	"net/http"
	"time"
)

type LocationService struct {
}

func (service LocationService) GetLocation(ip, location string) error {

	resp, err := http.Get("https://apis.map.qq.com/ws/geocoder/v1/?location=" + location + "&get_poi=1&key=GPNBZ-ZMBWP-GEPDJ-L56VX-TSLJ6-UAFDK")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	addr, err := ioutil.ReadAll(resp.Body)
	temp := model.Address{}

	temp.Ip = ip
	temp.Location = location
	temp.Addr = string(addr)
	temp.Createat = time.Now()

	_, err = DbEngin.Insert(temp)
	if err != nil {
		return err
	}
	return nil
}
