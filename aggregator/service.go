package aggregator

type ISrvProduct interface {
	CountAllProduct() int
	NameProduct() string
}

type ISrvProductType interface {
	CountAllProductType() int
	NameProductType() string
}
