package routers

import (
	"html/template"

	"github.com/jasonsoft/abb-ui/types"
	"github.com/jasonsoft/log"
	"github.com/jasonsoft/napnap"
)

func NewRouter() *napnap.Router {
	router := napnap.NewRouter()

	router.Get("/", loginEndpoint)

	// service
	router.Get("/services/:service_id", serviceDetailEndpoint)
	router.Get("/services/:service_id/logs", serviceLogsGetEndpoint)
	router.Get("/services", serviceListEndpoint)

	return router
}

func loginEndpoint(c *napnap.Context) {
	model := types.LoginModel{}
	model.PageData.APIHost = template.JS(_config.APIHost)

	err := c.Render(200, "login.html", model)
	if err != nil {
		log.Errorf("list render err: %v", err)
	}
}

func serviceListEndpoint(c *napnap.Context) {
	log.Debug("serviceListEndpoint")
	clusterName := c.Query("cluster")

	model := types.ServiceListPageData{
		ClusterID: clusterName,
	}
	model.PageData.APIHost = template.JS(_config.APIHost)

	err := c.Render(200, "service_list.html", model)
	if err != nil {
		log.Errorf("list render err: %v", err)
	}
}

func serviceDetailEndpoint(c *napnap.Context) {
	log.Debug("serviceDetailEndpoint")
	serviceID := c.Param("service_id")
	clusterName := c.Query("cluster")

	model := types.ServiceGetPageData{
		ClusterID: clusterName,
		ServiceID: serviceID,
	}
	model.PageData.APIHost = template.JS(_config.APIHost)

	err := c.Render(200, "service_detail.html", model)
	if err != nil {
		log.Error(err)
	}
}

func serviceLogsGetEndpoint(c *napnap.Context) {
	log.Debug("serviceLogsGetEndpoint")
	serviceID := c.Param("service_id")
	clusterID := c.Query("cluster")

	model := types.ServiceGetPageData{
		ClusterID: clusterID,
		ServiceID: serviceID,
	}
	model.PageData.APIHost = template.JS(_config.APIHost)

	err := c.Render(200, "service_logs.html", model)
	if err != nil {
		log.Error(err)
	}
}
