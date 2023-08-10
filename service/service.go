package service

import (
	"abc/model"
	"abc/repository"
	"errors"
)

type IServicePort interface {
	// port เอาไปให้ handler เรียกใช้
	GetSer() ([]model.GetRequestResponse, error)
	GetSerById(id string) (model.GetRequestResponse, error)
}

type serviceAdapter struct {
	// เก็บ port ของ repository เข้ามา เพื่อไปใช้งานต่อ
	r repository.IRepositoryPort
}

func NewServiceAdapter(r repository.IRepositoryPort) *serviceAdapter {
	// Adapter ไปเสียบ port ของ repository เข้ามา
	return &serviceAdapter{r: r}
}

// get all data
func (s serviceAdapter) GetSer() ([]model.GetRequestResponse, error) {
	beers, err := s.r.GetRep()
	if err != nil {
		err = errors.New("เชื่อมต่อฐานข้อมูลไม่สำเร็จ")
		return nil, err
	}

	var data []model.GetRequestResponse

	// แปลง database ให้เป็น json
	for _, beer := range beers {
		var a model.GetRequestResponse
		a.ID = beer.ID
		a.Name = beer.Name
		a.Type = beer.Type
		a.Detail = beer.Detail
		a.URL = beer.URL

		data = append(data, a)
	}
	return data, nil
}

// get data by id
func (s serviceAdapter) GetSerById(id string) (model.GetRequestResponse, error) {
	beer, err := s.r.GetRepById(id)
	if err != nil {
		err = errors.New("เชื่อมต่อฐานข้อมูลไม่สำเร็จ")
		return model.GetRequestResponse{}, err
	}

	var data model.GetRequestResponse
	data.ID = beer.ID
	data.Name = beer.Name
	data.Type = beer.Type
	data.Detail = beer.Detail
	data.URL = beer.URL

	return data, nil
}
