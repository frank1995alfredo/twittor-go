package bd

import (
	"context"
	"log"
	"time"

	"github.com/frank1995alfredo/twittor-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoTweets ... lee los tweets de un perfil
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultado []*models.DevuelvoTweets
	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find() //metodo de mongo
	opciones.SetLimit(20)      //cuando documentos me tiene que traer (paginando)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultado, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, false
		}
		resultado = append(resultado, &registro)
	}

	return resultado, true
}
