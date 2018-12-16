package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const DRIVER   = "mysql"
const USER     = "root"
const PASSWORD = "@root"
const DBNAME   = "test"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func main() {


	URL := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", USER,PASSWORD,DBNAME)

	conexao, erro := sql.Open(DRIVER, URL)

	if erro != nil {

		panic(erro.Error())

	}

	rs, erro := conexao.Query("select * from users")

	if erro != nil {

		panic(erro.Error())

	}

	var users []User

	for rs.Next() {

		var user User

		erro := rs.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

		if erro != nil {

			panic(erro.Error())

		}

		users = append(users, user)

	}

	for _,user := range users {

		fmt.Println(user.Name)

	}

	defer rs.Close()
	defer conexao.Close()


}