package inicializadores

import (
	"log"

	"github.com/joho/godotenv"
)

func CarregarVariaveisDeAmbiente() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}
