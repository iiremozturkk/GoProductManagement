package service

import (
	"GoProductManagement/models"
	"GoProductManagement/repository"
)

type ProductService interface {
	Create(product models.Product) (int, error)
	GetAll() ([]models.Product, error)
	GetByID(id int) (models.Product, error)
	Update(product models.Product) error
	Delete(id int) error
	UpdateStock(id int, quantity int) error // Stok güncelleme metodu
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(product models.Product) (int, error) {
	return s.repo.Create(product)
}

func (s *productService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetByID(id int) (models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) Update(product models.Product) error {
	return s.repo.Update(product)
}

func (s *productService) Delete(id int) error {
	return s.repo.Delete(id)
}


func (s *productService) UpdateStock(id int, quantity int) error {
	return s.repo.UpdateStock(id, quantity)
} 