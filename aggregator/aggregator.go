package aggregator

type ISrvAggregator interface {
	ISrvProduct
	ISrvProductType
}

type IRepoAggregator interface {
	IRepoProduct
	IRepoProductType
}

type Aggregator struct {
	SrvAggregator  ISrvAggregator
	RepoAggregator IRepoAggregator
}
