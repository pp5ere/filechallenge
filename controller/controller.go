package controller

import (
	"desafioNeoWay/entity"
	"desafioNeoWay/repository"
)

// Controllers contains the Controllers for each Entity.
type Controllers struct{
	Sales repository.SalesInterface
}

//Sales contains the injected Sales interface from Repository layer.
type Sales struct{
	Repository repository.SalesInterface
}

// New creates new Controllers for each Entity.
func New(repo *repository.DataBase) *Controllers {
	return &Controllers{
		Sales:  newSalesController(repo),	
	}
}

// SalesController contains methods that must be implemented by the injected layer.
type SalesController interface{
	Insert(s *entity.SalesData) error
}

func newSalesController(r *repository.DataBase) *Sales {
	return &Sales{
		Repository: r,
	}
}

//Insert requests the Repository layer to insert a new SalesData in the database.
func (s *Sales) Insert(sales *entity.SalesData) error {
	return s.Repository.Insert(sales)
}
