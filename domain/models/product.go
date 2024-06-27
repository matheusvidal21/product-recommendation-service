package models

type ProductDomain interface {
	GetID() string
	GetName() string
	GetPrice() float64
	GetCategory() CategoryDomain
}

type product struct {
	id       string
	name     string
	price    float64
	category CategoryDomain
}

func NewProductDomain(id, name string, price float64, category CategoryDomain) ProductDomain {
	return &product{
		id:       id,
		name:     name,
		price:    price,
		category: category,
	}
}

func (p *product) GetID() string {
	return p.id
}

func (p *product) GetName() string {
	return p.name
}

func (p *product) GetPrice() float64 {
	return p.price
}

func (p *product) GetCategory() CategoryDomain {
	return p.category
}
