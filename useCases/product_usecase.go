package usecases

import (
	"go-api/models"
	"go-api/repositories"
)

type ProductUseCase struct {
	repository repositories.ProductRepository
}

func NewProductUseCase(repo repositories.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product models.Product) (models.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return models.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(id int) (*models.Product, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
