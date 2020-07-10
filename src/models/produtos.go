package models

import "loja/db"

//Produto Struct
type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade, Id  int
}

//BuscaTodosProdutos func
func BuscaTodosProdutos() []Produto {
	db := db.ContectaDb()

	selecionaProdutos, err := db.Query("select * from produtos order by id asc;")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selecionaProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selecionaProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()

	return produtos
}

//CriarNovoProduto func
func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ContectaDb()

	insereDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

//DeletaProduto func
func DeletaProduto(id string) {
	db := db.ContectaDb()

	deletaProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletaProduto.Exec(id)
	defer db.Close()
}

//EditaProduto func
func EditaProduto(id string) Produto {
	db := db.ContectaDb()

	produtoBanco, err := db.Query("select * from produtos where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoAtualizacao := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoAtualizacao.Id = id
		produtoAtualizacao.Nome = nome
		produtoAtualizacao.Descricao = descricao
		produtoAtualizacao.Preco = preco
		produtoAtualizacao.Quantidade = quantidade
	}
	defer db.Close()
	return produtoAtualizacao
}

//AtualizarProduto func
func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ContectaDb()

	atualizaPrdouto, err := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}
	atualizaPrdouto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
