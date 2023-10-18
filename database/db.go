package database

import (
	"log"
	"os" //importando biblioteca
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	//os.Getenv("HOST") //biblioteca.pegando-variavel-de-ambiente("qual-var")
	//os.Getenv("USER")
	//os.Getenv("PASSWORD")
	//os.Getenv("DBNAME")
	//os.Getenv("PORT")
	stringDeConexao := "host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+" password="+os.Getenv("PASSWORD")+" dbname="+os.Getenv("DBNAME")+" port="+os.Getenv("PORT")+" sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
