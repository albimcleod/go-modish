package authentication

import (
	"fmt"
	"math/rand"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/*
//GenerateToken will create a JWT Token
func GenerateToken(siteid string) (string, error) {
	mySigningKey := []byte(MyVenueJwtSecret)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	claims["siteid"] = siteid
	token.Claims = claims

	return token.SignedString(mySigningKey)
}
*/

// ParseToken will parse the provided token for authentication
func ParseToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if token.Valid {
		return token, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, fmt.Errorf("That's not even a token")
		}
		return nil, fmt.Errorf("Couldn't handle this token:  %v", err)

	} else {
		return nil, fmt.Errorf("Couldn't handle this token: %v", err)
	}
}

//ValidatePassword will validate the password of a user
func ValidatePassword(userPass string, enteredPass string) bool {

	//p := GeneratePassword(enteredPass)
	//fmt.Printf("-----Password is: %v %v \n", enteredPass, p)

	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(enteredPass))

	return err == nil
}

//GeneratePassword will encrypt a string into a password
func GeneratePassword(password string) string {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(hashedPassword)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandomString will create a random string
func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
