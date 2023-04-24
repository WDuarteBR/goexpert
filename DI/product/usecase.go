package product

type ProductUseCase struct {
	repository *ProductRepository
}

func NewProductUseCase(repo *ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repository: repo,
	}
}

// Este método retorna uma entidade Product
// onde o correto é retornar um DTO
// mas para simplicifar manterei assim
func (p *ProductUseCase) GetProduct(id int) (Product, error) {
	return p.repository.GetProduct(id)
}
