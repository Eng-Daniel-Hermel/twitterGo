package models

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	Email string             `json:"email`
	ID    primitive.ObjectID `bson:"id" json:"_id,omitempy"`
	jwt.RegisteredClaims
}
