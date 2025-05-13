package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(newProduct model.Product, id_product int) (bool, error) {
	newProduct.ID = id_product
	rowsAffected, err := pu.repository.UpdateProduct(newProduct)
	if err != nil {
		return false, err
	}

	return rowsAffected, nil
}

func (pu *ProductUsecase) DeleteProduct(id_product int) (bool, error) {
	rowsAffected, err := pu.repository.DeleteProduct(id_product)
	if err != nil {
		return false, err
	}

	return rowsAffected, nil
}
