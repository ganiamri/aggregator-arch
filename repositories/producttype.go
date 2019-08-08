package repositories

import (
	nmodels "ketitik/ganiamri/asset-mgmt/models"
)

type RepoProductType struct {
}

func (r *RepoProductType) GetAllProductTypes() []nmodels.ProductType {
	return []nmodels.ProductType{
		nmodels.ProductType{
			ID:   1,
			Name: "Product Type 1",
		},
		nmodels.ProductType{
			ID:   2,
			Name: "Product Type 2",
		},
	}
}

func (r *RepoProductType) GetProductType(id int) *nmodels.ProductType {
	return &nmodels.ProductType{
		ID:   1,
		Name: "Product Type 1",
	}
}
