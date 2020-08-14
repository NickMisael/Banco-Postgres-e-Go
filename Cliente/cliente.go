package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func limpaTela() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

const (
	host   = "localhost"
	port   = 5432
	user   = "user0"
	dbname = "Banco"
)

//Struct Conta
type Conta struct {
	NumConta  int
	FirstName string
	LastName  string
	Idade     int
	Email     string
	Saldo     float32
}

func Cadastrar(dbparam *sql.DB) {
	c := Conta{}
	fmt.Printf("Digite o Primeiro Nome: ")
	fmt.Scanf("%s", &c.FirstName)
	fmt.Printf("Digite o Sobrenome: ")
	fmt.Scanf("%s", &c.LastName)
	fmt.Printf("Digite a Idade: ")
	fmt.Scanf("%d", &c.Idade)
	fmt.Printf("Digite o Email: ")
	fmt.Scanf("%s", &c.Idade)
	limpaTela()
	fmt.Println("Carregando Dados...")
	time.Sleep(time.Second + 3)

	// Inserindo dados
	sqlInsert := `
	INSERT INTO users (age, first_name, last_name, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	err := dbparam.QueryRow(sqlInsert, c.Idade, c.FirstName, c.LastName, c.Email).Scan(&c.NumConta)
	if err != nil {
		panic(err)
	}
	fmt.Println("O Novo registro é ->", c.NumConta)
	time.Sleep(time.Second + 3)
	limpaTela()
}

func Consultar(dbparam *sql.DB) {
	limpaTela()
	rows, err := dbparam.Query("SELECT id, first_name, last_name FROM users LIMIT $1", 10)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		c := Conta{}
		err = rows.Scan(&c.NumConta, &c.FirstName, &c.LastName)
		if err != nil {
			panic(err)
		}
		if c.FirstName == "" && c.NumConta == 1 {
			fmt.Println("Nenhum dado foi encontrado!")
		} else {
			fmt.Println("ID ->", c.NumConta)
			fmt.Println("Nome ->", c.FirstName, c.LastName)
			fmt.Println("")
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func Deletar(dbParam *sql.DB) {
	limpaTela()
	var id int
	var err error
	for {
		var i string
		fmt.Printf("Digite o ID a ser removido: ")
		fmt.Scanf("%s", &i)
		if id, err = strconv.Atoi(i); err != nil {
			fmt.Println("Digite um número válido")
		} else if id < 1 {
			fmt.Println("Digite um número válido")
		} else {
			break
		}
	}
	limpaTela()
	// DELETANDO
	sqlDelete := `
	DELETE FROM users
	WHERE id = $1`
	res, er := dbParam.Exec(sqlDelete, id)
	if er != nil {
		panic(er)
	}
	count, er := res.RowsAffected()
	if er != nil {
		panic(er)
	}
	fmt.Println(count)
	fmt.Println("Dados exluídos com sucesso!")
	time.Sleep(time.Second + 3)
}

func Open(dbsource string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", dbsource)
	if err != nil {
		panic(err)
	}

	/*err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexão bem-sucedida")*/
	return
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := Open(psqlInfo)
	if err != nil {
		panic(err)
	}

	for {
		var esc string
		limpaTela()
		print("1 - Adicionar Cliente")
		print("2 - Alterar Cliente")
		print("3 - Excluir Cliente")
		print("4 - Buscar Cliente")
		print("5 - Listar Cliente")
		print("6 - Sair do sistema Banco")
		fmt.Printf("-> ")
		fmt.Scanf("%s", &esc)

		if es, er := strconv.Atoi(esc); er != nil {
			fmt.Println("Engraçadão você hein!!")
			fmt.Println("Digite um número válido!!")
			time.Sleep(time.Second + 4)
		} else if es < 1 || es > 6 {
			fmt.Println("Engraçadão você hein!!")
			fmt.Println("Digite um número válido!!")
			time.Sleep(time.Second + 4)
		} else {
			if es == 6 {
				fmt.Println("Obrigado por utilizar!!")
				break
			} else {
				switch es {
				case 1:
					limpaTela()
					Cadastrar(db)
				case 2:
					limpaTela()
					Consultar(db)
					for {
						fmt.Scanf("%s", &esc)
						if esc == "0" {
							break
						}
					}
					time.Sleep(time.Second + 2)
					limpaTela()
				case 3:
					Deletar(db)
				}
			}
		}
	}
}
