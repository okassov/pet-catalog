package usecase

import (
	"context"

	"github.com/okassov/pet-catalog/internal/entity"
)

type CategoryUseCase struct {
	repo CategoryRepo
}

func NewCategoryUseCase(r CategoryRepo) *CategoryUseCase {
	return &CategoryUseCase{
		repo: r,
	}
}

func (uc *CategoryUseCase) GetCategory(ctx context.Context, p entity.Category) error {
	return nil
}

func (uc *CategoryUseCase) CreateCategory(ctx context.Context, p entity.Category) error {
	return nil
}

func (uc *CategoryUseCase) UpdateCategory(ctx context.Context, p entity.Category) error {
	return nil
}

func (uc *CategoryUseCase) DeleteCategory(ctx context.Context, p entity.Category) error {
	return nil
}

func (uc *CategoryUseCase) ListCategories(ctx context.Context, p entity.Category) error {
	return nil
}
