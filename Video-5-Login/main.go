package main

import(
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {


	senha := "s3cr3t"

	cost := bcrypt.DefaultCost

	hash, erro := bcrypt.GenerateFromPassword([]byte(senha), cost)

	if erro != nil {

		panic(erro.Error())

	}

	//fmt.Println(string(hash))
	
	var inputSenha string

	fmt.Print("Entre com sua senha: ")
	fmt.Scan(&inputSenha)

	erro = bcrypt.CompareHashAndPassword(hash, []byte(inputSenha))

	if erro != nil {

		fmt.Println("Senha incorreta!")
		return

	}

	fmt.Println("Logado com sucesso!")

}