package actio

import (
	"axo/data"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret is the secret key for the JWT token
var JWTSecret = "cubidronun şifresi ç0k g1zl1"

// JwtToken creates a new token with the user id
func JwtToken(user data.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"ID":       user.ID,
			"PASSWORD": user.Password,
			"MAIL":     user.Mail,
			"NAME":     user.Name,
			"IMAGE":    user.Image,
			"PERMS":    user.Perms,
			"VERIFIED": user.Verified,
			"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		})

	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JwtVerify checks the token and returns the user id
func JwtVerify(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", nil
	}
	var id string = token.Claims.(jwt.MapClaims)["ID"].(string)
	return id, nil
}
