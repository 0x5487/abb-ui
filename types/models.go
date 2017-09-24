package types

import "html/template"

type PageData struct {
}

type ServiceListPageData struct {
	PageData
	ClusterID string
	APIHost   template.JS
}

type ServiceGetPageData struct {
	PageData
	ClusterID string
	ServiceID string
	APIHost   template.JS
}
