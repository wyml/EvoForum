package jwt

import (
	"errors"
	"forum/pkg/app"
	"forum/pkg/config"
	"forum/pkg/logger"
	"github.com/gin-gonic/gin"
	"strings"
	"time"

	jwtpkg "github.com/golang-jwt/jwt"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct {
	// 密钥
	SignKey []byte

	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration
}

type JwtCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime string `json:"expire_time"`

	jwtpkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

func (j *JWT) ParserToken(c *gin.Context) (*JwtCustomClaims, error) {

	tokenString, parseErr := j.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	token, err := j.parseTokenString(tokenString)

	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrHeaderMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

func (j *JWT) RefreshToken(c *gin.Context) (string, error) {
	tokenString, parseErr := j.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	token, err := j.parseTokenString(tokenString)

	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}

	claims := token.Claims.(*JwtCustomClaims)

	x := app.TimenowInTimezone().Add(-j.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		claims.StandardClaims.ExpiresAt = j.expireAtTime()
		return j.createToken(*claims)
	}
}

func (j *JWT) IssueToken(userID string, userName string) string {

	expireAtTime := j.expireAtTime()
	claims := JwtCustomClaims{
		userID,
		userName,
		string(expireAtTime),
		jwtpkg.StandardClaims{
			NotBefore: app.TimenowInTimezone().Unix(),
			IssuedAt:  app.TimenowInTimezone().Unix(),
			ExpiresAt: expireAtTime,
			Issuer:    config.GetString("app.name"),
		},
	}

	token, err := j.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

func (j *JWT) createToken(claims JwtCustomClaims) (string, error) {
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

func (j *JWT) expireAtTime() int64 {
	timenow := app.TimenowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	return timenow.Add(expire).Unix()
}

func (j *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return j.SignKey, nil
	})
}

func (j *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}
