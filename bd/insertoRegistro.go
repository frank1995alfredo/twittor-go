package bd

import (
	"context"
	"time"

	"github.com/frank1995alfredo/twittor-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoRegistro ... es la parada final con BD para insertar los datos del usuario
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	//para que no supere los 15 seg
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
