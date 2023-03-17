package repository

import (
	"context"

	"github.com/okassov/pet-catalog/internal/entity"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	gorm.DB
}

func NewCategoryRepo(pg gorm.DB) *ProductRepo {
	return &ProductRepo{&pg}
}

func (r *ProductRepo) GetCategory(ctx context.Context, p entity.Category) error {
	return nil
}

func (r *ProductRepo) CreateCategory(ctx context.Context, p entity.Category) error {
	return nil
}

func (r *ProductRepo) UpdateCategory(ctx context.Context, p entity.Category) error {
	return nil
}

func (r *ProductRepo) DeleteCategory(ctx context.Context, p entity.Category) error {
	return nil
}

func (r *ProductRepo) ListCategories(ctx context.Context, p entity.Category) error {
	return nil
}
