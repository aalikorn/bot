package product

type Service struct{}

func newService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}
