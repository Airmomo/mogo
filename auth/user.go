package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"mogo/util"
)

type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GetUserToken(id uint) (string, error) {
	c := UserClaims{
		UserID:         id,
		StandardClaims: util.GetJwtStandardClaims(),
	}
	//使用指定的签名方法创建JWT签名对象
	token := jwt.NewWithClaims(util.GetSigningMethod(), c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(util.GetSecret())
}

// ParseUserToken UserClaims解析JWT的方法
func ParseUserToken(tokenString string) (*UserClaims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return util.GetSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	// 校验令牌
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
