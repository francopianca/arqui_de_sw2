package services

import (
	"github.com/francopianca/arqui_de_sw2/dtos"
	e "github.com/francopianca/arqui_de_sw2/utils/errors"
)

type Service interface {
	Get(id string) (dtos.ItemDTO, e.ApiError)
	Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError)
}
