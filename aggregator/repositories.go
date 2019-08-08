package aggregator

import (
	nmodels "ketitik/ganiamri/asset-mgmt/models"
)

type IRepoProduct interface {
	GetAllProducts() []nmodels.Product
	GetProduct(id int) *nmodels.Product
}

type IRepoProductType interface {
	GetAllProductTypes() []nmodels.ProductType
	GetProductType(id int) *nmodels.ProductType
}
