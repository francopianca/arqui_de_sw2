package repositories

import (
	"github.com/francopianca/arqui_de_sw2/dtos"
	"github.com/francopianca/arqui_de_sw2/utils/errors"
)

type Repository interface {
	Get(id string) (dtos.ItemDTO, errors.ApiError)
	Insert(item dtos.ItemDTO) (dtos.ItemDTO, errors.ApiError)
	Update(item dtos.ItemDTO) (dtos.ItemDTO, errors.ApiError)
	Delete(id string) errors.ApiError
}
