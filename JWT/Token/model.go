package Token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	Department string `json:"department"`
	jwt.StandardClaims
}

const TokenExpireDurnTime = time.Hour

var MySecret = []byte("are you ready")
