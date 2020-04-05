package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	//initTokenExpireDuration 过期时间
	initTokenExpireDuration time.Duration
	//initSecret JWT密钥
	initSecret string
	//initSecret JWT的签发人
	initIssuer string
	//initSigningMethod JWT的声明加密算法
	initSigningMethod *jwt.SigningMethodHMAC
)

// BuildJwtStandardClaims 初始化Jwt.StandardClaims属性
func BuildJwtStandardClaims(tokenExpireDuration time.Duration, secret string, issuer string, signingMethod string) {
	initTokenExpireDuration = tokenExpireDuration
	initSecret = secret
	initIssuer = issuer
	switch signingMethod {
	case "HS512":
		initSigningMethod = jwt.SigningMethodHS512
	case "HS384":
		initSigningMethod = jwt.SigningMethodHS384
	case "HS256":
		initSigningMethod = jwt.SigningMethodHS256
	default:
		initSigningMethod = jwt.SigningMethodHS256
	}
	if initTokenExpireDuration == 0 {
		initTokenExpireDuration = time.Hour * 2
	}
	if initSecret == "" {
		initSecret = "momo爱爱爱da"
	}
	if initIssuer == "" {
		initIssuer = "myProject"
	}
}

func GetJwtStandardClaims() jwt.StandardClaims {
	//声明jwt注册的标准 (建议但不强制使用)
	return jwt.StandardClaims{
		//Audience:  "",// 接收jwt的一方
		ExpiresAt: time.Now().Add(initTokenExpireDuration).Unix(), // jwt的过期时间，过期时间必须要大于签发时间
		//Id:        "",// jwt的唯一身份标识，主要用来作为一次性token，从而回避重放攻击
		IssuedAt: time.Now().Unix(), // 签发时间
		Issuer:   initIssuer,        // jwt签发者
		//NotBefore: 0,// 定义在什么时间之前，该jwt都是不可用的
		//Subject:   "",// jwt所面向的用户
	}
}

func GetSecret() []byte {
	return []byte(initSecret)
}

func GetSigningMethod() *jwt.SigningMethodHMAC {
	return initSigningMethod
}

func GetExpiresTime() int64 {
	return time.Now().Add(initTokenExpireDuration).Unix()
}
