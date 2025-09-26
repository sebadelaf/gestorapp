// Archivo de configuracion de la base de datos MongoDB donde se realiza la conexion a la base de datos
package models

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Variables de entorno para la conexion a la base de datos, al igual que en c solo se les reconoce su tipo
// var nombre tipo
// atajo a definir x:= valor
var URI string
var PORT string
var DB_NAME string
var COLLECTION_NAME string
var mongoClient *mongo.Client //este mantiene la conexion a la base de datos

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	URI = os.Getenv("MONGODB_URI")
	PORT = os.Getenv("PORT")
	DB_NAME = os.Getenv("DB_NAME")
	COLLECTION_NAME = os.Getenv("COLLECTION_NAME")
}

func ConnectDb() {
	LoadEnv()
	clientOption := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	mongoClient = client
	log.Println("Connected to MongoDB!")
}
