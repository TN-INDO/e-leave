package helpers

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego/config"
	"github.com/dgrijalva/jwt-go"
)

// EzToken ...
type EzToken struct {
	ID      int64
	Email   string
	Expires int64
}

var verifyKey string

func init() {
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	verifyKey = appConf.String("jwt::token")
}

// GetToken ...
func (e *EzToken) GetToken() (string, error) {
	var strID = strconv.FormatInt(e.ID, 8)

	claims := &jwt.StandardClaims{
		ExpiresAt: e.Expires,
		Id:        strID,
		Audience:  e.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(verifyKey))
	if err != nil {
		CheckErr("Error signed string @GetToken", err)
		panic(err)
	}

	return tokenString, err
}

// ValidateToken ...
func (e *EzToken) ValidateToken(tokenString string) (bool, error) {
	if tokenString == "" {
		return false, errors.New("No Token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})

	if token == nil {
		return false, errors.New("Token Invalid")
	}

	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("Invalid Token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("Token Expired")
		} else {
			return false, errors.New("Other Error")
		}
	} else {
		return false, errors.New("Other Error")
	}
}
