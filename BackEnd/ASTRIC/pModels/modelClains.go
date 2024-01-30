package pmodels

import (
	"github.com/dgrijalva/jwt-go"
)

// Clain modelo de desencriptado del tocken
type Clain struct {
	Hash string `json:"hash"`
	ID   int    `json:"id"`
	jwt.StandardClaims
}
