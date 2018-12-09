package main

import(
	"fmt"
	"./models"
)

func main() {

	ListProducts()

	//CreateProduct()
	//UpdateProduct()
	//DeleteProduct()

	Search()

	
}

func ListProducts() {

	products, erro := models.GetAll()

	if erro != nil {

		fmt.Println("Oops! Ocorreu um erro ao consultar a tabela de produtos")
		return 

	}

	fmt.Println("\n\nLista de produtos:")

	for _, product := range products {

		fmt.Println(product)

	}

	fmt.Println("\n\n")
}

func CreateProduct() {

	var product models.Product

	fmt.Print("Descricao: ")
	fmt.Scan(&product.Description)

	fmt.Print("Quantidade: ")
	fmt.Scan(&product.Quantity)

	fmt.Print("Preco: ")
	fmt.Scan(&product.Price)

	product.Amount = product.Price * float64(product.Quantity)

	_, erro := models.NewProduct(product)

	if erro != nil {

		fmt.Println("Oops! Ocorreu um erro ao inserir!")
		return

	} 

	fmt.Println("Produto adicionado com sucesso!")

}

func UpdateProduct() {

	var product models.Product

	fmt.Print("Descricao: ")
	fmt.Scan(&product.Description)

	fmt.Print("Quantidade: ")
	fmt.Scan(&product.Quantity)

	fmt.Print("Preco: ")
	fmt.Scan(&product.Price)

	fmt.Print("ID: ")
	fmt.Scan(&product.Id)

	product.Amount = product.Price * float64(product.Quantity)

	_, erro := models.UpdateProduct(product)

	if erro != nil {

		fmt.Println("Oops! Ocorreu um erro ao atualizar!")
		return

	} 

	fmt.Println("Produto atualizado com sucesso!")

}

func DeleteProduct() {

	var id string

	fmt.Print("ID: ")
	fmt.Scan(&id)

	rows, erro := models.DeleteProduct( "id", id )

	if erro != nil {

		fmt.Println("Oops! Ocorreu um erro ao deletar!")
		return

	}

	fmt.Printf("%d linhas deletadas.", rows)

}

func Search() {

	var id int

	fmt.Print("Buscar por ID: ")
	fmt.Scan(&id)

	product, erro := models.Find(id)
	
	if erro != nil {

		fmt.Println("Oops! Ocorreu um erro interno!")
		return

	}

	if product.Id != 0 {

		fmt.Println("\n\nProduto encontrado:")
		fmt.Println(product)
		fmt.Println("\n")

		return
		
	}

	fmt.Println("\n\nProduto NÃO encontrado...")
}