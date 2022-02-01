package configs

import(
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro loading .env file")
	}

	return os.Getenv("MONGOURI")
}