package auth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// TODO: hide key
var secretKey = []byte("secret-key")

func CreateToken(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "email": email, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

type JWTClaims struct {
    Username string `json:"username"`
    jwt.Claims
}


func VerifyToken(tokenString string) (map[string]interface{}, error) {
    token, err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error) {
       return secretKey, nil
    })
   
    if err != nil {
       return nil, err
    }
   
    if !token.Valid {
       return nil, fmt.Errorf("invalid token")
    }
   
    return token.Claims.(jwt.MapClaims), nil
}
