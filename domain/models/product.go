package models

type ProductDomain interface {
	GetID() string
	GetName() string
	GetPrice() float64
}

type product struct {
	id    string
	name  string
	price float64
}

func NewProductDomain(id, name string, price float64) ProductDomain {
	return &product{
		id:    id,
		name:  name,
		price: price,
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
