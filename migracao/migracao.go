package main

import (
	"projeto/inicializadores"
	"projeto/modelos"
)

func init() {
	inicializadores.CarregarVariaveisDeAmbiente()
	inicializadores.ConectarAoBD()
}

func main() {
	inicializadores.BD.AutoMigrate(&modelos.Post{})
}
