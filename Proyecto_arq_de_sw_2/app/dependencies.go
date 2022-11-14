package app

import (
	"github.com/francopianca/arqui_de_sw2/tree/main/ej-books/controllers"
	service "github.com/francopianca/arqui_de_sw2/tree/main/ej-books/services"
	"github.com/francopianca/arqui_de_sw2/tree/main/ej-books/services/repositories"
	"time"
)

type Dependencies struct {
	BookController *controllers.Controller
}

func BuildDependencies() *Dependencies {
	// Repositories
	ccache := repositories.NewCCache(1000, 100, 30*time.Second)
	memcached := repositories.NewMemcached("localhost", 11211)
	mongo := repositories.NewMongoDB("localhost", 27017, "items")

	// Services
	service := service.NewServiceImpl(ccache, memcached, mongo)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		ItemController: controller,
	}
}
