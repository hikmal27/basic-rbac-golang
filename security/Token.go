package security

import (
	"fmt"
	"gorbac/app/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user entity.User) (string, error) {
	var mySigningKey = []byte("98hbun98h")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
