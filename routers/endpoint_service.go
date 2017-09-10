package routers

import (
	"html/template"

	"github.com/jasonsoft/log"
	"github.com/jasonsoft/napnap"
)

func NewRouter() *napnap.Router {
	router := napnap.NewRouter()

	// service
	router.Get("/services", serviceListEndpoint)

	return router
}

type ApiConfig struct {
	APIHost template.JS
}

func serviceListEndpoint(c *napnap.Context) {
	apiConfig := ApiConfig{
		APIHost: template.JS(_config.APIHost),
	}
	log.Infof("api: %s", apiConfig.APIHost)
	c.Render(200, "service_index.html", apiConfig)
}
