package repositories

import (
	"fmt"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	e "github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
	"github.com/karlseguin/ccache/v2"
	"time"
)

type RepositoryCCache struct {
	Client     *ccache.Cache
	DefaultTTL time.Duration
}

func NewCCache(maxSize int64, itemsToPrune uint32, defaultTTL time.Duration) *RepositoryCCache {
	client := ccache.New(ccache.Configure().MaxSize(maxSize).ItemsToPrune(itemsToPrune))
	fmt.Println("[CCache] Initialized")
	return &RepositoryCCache{
		Client:     client,
		DefaultTTL: defaultTTL,
	}
}

func (repo *RepositoryCCache) Get(id string) (dtos.ItemDTO, e.ApiError) {
	item := repo.Client.Get(id)
	if item == nil {
		return dtos.ItemDTO{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	if item.Expired() {
		return dtos.Item{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	return item.Value().(dtos.ItemDTO), nil
}

func (repo *RepositoryCCache) Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	repo.Client.Set(item.Id, item, repo.DefaultTTL)
	return item, nil
}

func (repo *RepositoryCCache) Update(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	repo.Client.Set(item.Id, item, repo.DefaultTTL)
	return item, nil
}

func (repo *RepositoryCCache) Delete(id string) e.ApiError {
	repo.Client.Delete(id)
	return nil
}
