package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN ... variable export
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:1234@cluster0-pbpgn.gcp.mongodb.net/test?retryWrites=true&w=majority")

//ConectarBD ... es la funcion que me permite conectar la base de datos
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil) //ping a la base de datos para ver si esta conectado
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

//ChequeoConnection ... chequeo a la conexion
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
