package services

import (
	"REST-GO-CLEAN/lib"

	"REST-GO-CLEAN/api/repository"
)

type ProductService struct {
	repo   repository.ProductRepository
	env    lib.Env
	logger lib.Logger
}

func NewProductService(
	repo repository.ProductRepository,
	env lib.Env,
	logger lib.Logger,
) ProductService {
	return ProductService{
		repo:   repo,
		env:    env,
		logger: logger,
	}
}
