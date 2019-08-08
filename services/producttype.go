package services

import (
	nagg "ketitik/ganiamri/asset-mgmt/aggregator"
)

type SrvProductType struct {
	nagg.ISrvAggregator
}

func (s *SrvProductType) CountAllProductType() int {
	return 0
}

func (s *SrvProductType) NameProductType() string {
	return "Product Type Name"
}
