package model

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Products struct {
	Product []Product
}

func (p *Products) Add(product Product) {
	p.Product = append(p.Product, &product)
}

func NewProduct() *Product {
	return &Product
}
