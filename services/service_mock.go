package services

import (
	"github.com/francopianca/arqui_de_sw2/dtos"
	e "github.com/francopianca/arqui_de_sw2/utils/errors"
)

type ServiceMock struct{}

func NewServiceMock() ServiceMock {
	return ServiceMock{}
}

func (ServiceMock) Get(id string) (dtos.ItemDTO, e.ApiError) {
	return dtos.ItemDTO{
		Id:     "12345",
		Titulo: "Mocked item",
	}, nil
}

func (ServiceMock) Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	return dtos.ItemDTO{
		Id:     "12345",
		Titulo: item.Titulo,
	}, nil
}
