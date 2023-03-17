package usecase

import (
	"context"

	"github.com/okassov/pet-catalog/internal/entity"
)

type ProductUseCase struct {
	repo ProductRepo
}

func NewProductUseCase(r ProductRepo) *ProductUseCase {
	return &ProductUseCase{
		repo: r,
	}
}

func (uc *ProductUseCase) GetProduct(ctx context.Context, p entity.Product) error {
	return nil
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, p entity.Product) error {
	return nil
}

func (uc *ProductUseCase) UpdateProduct(ctx context.Context, p entity.Product) error {
	return nil
}

func (uc *ProductUseCase) DeleteProduct(ctx context.Context, p entity.Product) error {
	return nil
}

func (uc *ProductUseCase) ListProducts(ctx context.Context, p entity.Product) error {
	return nil
}
