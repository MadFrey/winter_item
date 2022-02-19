package service

import (
	"blog/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)
var accessSecret = []byte("hyhjzy")
var refreshSecret = []byte("ar")

func CreateToken(username string) (string,string ){
	claims:=model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*3,
			Issuer:    "Alsace",
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenSigned,err:=token.SignedString(accessSecret)
	if err != nil {
		panic(err)
		return "",""
	}
	rT:=model.MyClaims{
		Username: username,
		StandardClaims:jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt:time.Now().Add(time.Hour).Unix(),
			Issuer:    "Alsace",
		},
	}
	refreshToken :=jwt.NewWithClaims(jwt.SigningMethodHS256,rT)
	refreshTokenSigned,err:= refreshToken.SignedString(refreshSecret)
	if err != nil {
		panic(err)
		return "",""
	}
	return tokenSigned,refreshTokenSigned
}

func ParseToken(tokenString string,refreshTokenString string) (*model.MyClaims, bool, error) {
	token,err:= jwt.ParseWithClaims(tokenString,&model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret,nil
	})
	if claims,ok:=token.Claims.(*model.MyClaims);ok&&token.Valid {
		return claims,false,nil
	}
	refreshToken,err:= jwt.ParseWithClaims(refreshTokenString,&model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret,nil
	})
	if err != nil {
		return nil, false, err
	}
	if claims,ok:=refreshToken.Claims.(*model.MyClaims);ok&&refreshToken.Valid {
		return claims,true,nil
	}
	return nil,false,errors.New("invalid token")
}


