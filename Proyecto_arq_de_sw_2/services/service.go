package services

import (
	"github.com/francopianca/arqui_de_sw2/tree/main/ej-books/dtos"
	e "github.com/francopianca/arqui_de_sw2/tree/main/ej-books/utils/errors"
)

type Service interface {
	Get(id string) (dtos.ItemDTO, e.ApiError)
	Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError)
}
