package handlers

import (
	nagg "ketitik/ganiamri/asset-mgmt/aggregator"
)

type IHandlerProduct interface {
	GetCountAndNameProduct(id int) string
}

type HandlerProduct struct {
	nagg.ISrvAggregator
}

func (h *HandlerProduct) GetCountAndNameProduct(id int) string {
	return ""
}
