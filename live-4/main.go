package main

func main() {
	product1 := model.NewProduct()
	product1.Name = "Carrinho"

	product2 := model.NewProduct()
	product1.Name = "boneca"

	products := model.Products{}
	products.Add(product1)
	products.Add(product2)
}
