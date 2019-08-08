package services

import (
	nagg "ketitik/ganiamri/asset-mgmt/aggregator"
)

type SrvProduct struct {
	nagg.IRepoAggregator
}

func (s *SrvProduct) CountAllProduct() int {
	return 0
}

func (s *SrvProduct) NameProduct() string {
	return "Product Name"
}
