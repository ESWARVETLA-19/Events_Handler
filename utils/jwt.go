package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secreatKey="supersecreat"


func GenerateToken(email string,userId int64)(string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	})
	return token.SignedString([]byte(secreatKey))
}

func VarifyToken(token string)error{
	parsedToken,err:=jwt.Parse(token,func(token *jwt.Token)(interface{},error){
		_,status:=token.Method.(*jwt.SigningMethodHMAC)
		if !status{
			return nil,errors.New("unexpected signing method")
		}
		return secreatKey,nil
	})
	if err!=nil{
		return errors.New("could not parseToken")
	}
	tokenisValid:=parsedToken.Valid
	if !tokenisValid{
		return errors.New("Invalid Token")
	}
	// claims,ok:=parsedToken.Claims.(*jwt.MapClaims)
	// if !ok{
	// 	return errors.New("invalid token claims")
	// }
	// email:=claims["email"].(string)
	// userId:=claims["userId"].(int64)
	return nil
}