package handlers

import (
	nagg "ketitik/ganiamri/asset-mgmt/aggregator"
)

type IHandlerProductType interface {
	GetCountAndNameProductType(id int) string
}

type HandlerProductType struct {
	nagg.ISrvAggregator
}

func (h *HandlerProductType) GetCountAndNameProductType(id int) string {
	return ""
}
