package usecase

import (
	"context"
	"errors"
	"product-microservice/product"
	"time"

	"product-microservice/domain"
)

type productUsecase struct {
	productRepo product.Repository
}

func NewProductUsecase(productRepo product.Repository) product.Usecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (p *productUsecase) CreateProduct(ctx context.Context, product domain.Product) error {
	err := p.productRepo.CreateProduct(ctx, product)

	if err != nil {
		return err
	}
	return nil
}

func (p *productUsecase) ListProducts(c context.Context) (interface{}, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*2)
	defer cancel()

	products := p.productRepo.ListProducts(ctx)

	select {
	case <-ctx.Done():
		return nil, errors.New("request timeout")
	default:
	}

	return products, nil
}
