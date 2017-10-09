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
