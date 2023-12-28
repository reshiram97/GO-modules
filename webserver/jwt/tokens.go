package jwt

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
);

var secretKey = []byte(os.Getenv("SECRET_KEY"))
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

func VerifyToken(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return secretKey, nil
	})
	if err != nil {
	   return err
	}
	if !token.Valid {
	   return nil
	} 
	return nil
}

func ExtractClaims(tokenString string) (string,error) {
	if tokenString == "" {
		return "Invalid",nil
	}
	tokenString = tokenString[len("Bearer "):]
	// log.Fatalf(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey,nil
	})
	if err != nil {
		return "Invalid",nil
	 }
	if !token.Valid {
		return "Invalid",nil
	} 
	claims,ok:= token.Claims.(jwt.MapClaims)
	if ok {
		email:= claims["email"].(string)
		return email,nil
	}
	return "invalid",nil
	
}