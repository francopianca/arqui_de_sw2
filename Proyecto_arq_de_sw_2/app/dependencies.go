package app

import (
	"time"

	service "github.com/francopianca/arqui_de_sw2/tree/main/Proyecto_arq_de_sw_2//services"
	"github.com/francopianca/arqui_de_sw2/tree/main/Proyecto_arq_de_sw_2/controllers"
	"github.com/francopianca/arqui_de_sw2/tree/main/Proyecto_arq_de_sw_2/services/repositories"
)

type Dependencies struct {
	ItemController *controllers.Controller
}

func BuildDependencies() *Dependencies {
	// Repositories
	ccache := repositories.NewCCache(1000, 100, 30*time.Second)
	memcached := repositories.NewMemcached("localhost", 11211)
	mongo := repositories.NewMongoDB("localhost", 27017, "avisos")
	solr := repositories.NewSolrClient("localhost", 8983, "avisos")

	// Services
	service := service.NewServiceImpl(ccache, memcached, mongo, solr)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		ItemController: controller,
	}
}
