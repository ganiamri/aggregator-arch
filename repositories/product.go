package repositories

import (
	nmodels "ketitik/ganiamri/asset-mgmt/models"
)

type RepoProduct struct {
}

func (r *RepoProduct) GetAllProducts() []nmodels.Product {
	return []nmodels.Product{
		nmodels.Product{
			ID:            1,
			ProductTypeID: 1,
			Name:          "Product 1",
			Quantity:      1,
		},
		nmodels.Product{
			ID:            2,
			ProductTypeID: 2,
			Name:          "Product 2",
			Quantity:      2,
		},
	}
}

func (r *RepoProduct) GetProduct(id int) *nmodels.Product {
	return &nmodels.Product{
		ID:            1,
		ProductTypeID: 1,
		Name:          "Product 1",
		Quantity:      1,
	}
}
