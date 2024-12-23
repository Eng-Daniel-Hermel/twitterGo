package bd

import (
	"context"

	"github.com/Eng-Daniel-Hermel/twitterGo.git/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegisstro(u models.Usuario) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
