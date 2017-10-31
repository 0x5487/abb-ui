package types

import "html/template"

type PageData struct {
	APIHost template.JS
}

type ServiceListPageData struct {
	PageData
	ClusterID string
}

type ServiceGetPageData struct {
	PageData
	ClusterID string
	ServiceID string
}

type LoginModel struct {
	PageData
}

type ConfigPageData struct {
	PageData
	ClusterID string
	ConfigID  string
}

type ConfigGetPageData struct {
	PageData
	ClusterID string
	ConfigID  string
}
