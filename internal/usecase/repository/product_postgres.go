package repository

import (
	"context"

	"github.com/okassov/pet-catalog/internal/entity"
	"gorm.io/gorm"
)

type ProductRepo struct {
	*gorm.DB
}

func NewProductRepo(pg gorm.DB) *ProductRepo {
	return &ProductRepo{&pg}
}

func (r *ProductRepo) GetProduct(ctx context.Context, p entity.Product) error {
	if result := r.DB.Find(p, p.ID); result.Error != nil {
		return result.Error
	}
	// query, args, err := r.Builder.
	// 	Select("*").
	// 	From("users").
	// 	Where("username IN (?) AND email IN (?)", a.Username, a.Email).
	// 	ToSql()

	// if err != nil {
	// 	return nil, fmt.Errorf("UserRepo - Get - r.Builder: %w", err)
	// }

	return nil
}

func (r *ProductRepo) CreateProduct(ctx context.Context, p entity.Product) error {
	return nil
}

func (r *ProductRepo) UpdateProduct(ctx context.Context, p entity.Product) error {
	return nil
}

func (r *ProductRepo) DeleteProduct(ctx context.Context, p entity.Product) error {
	return nil
}

func (r *ProductRepo) ListProducts(ctx context.Context, p entity.Product) error {
	return nil
}
