package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"zero-shop/common/ctxData"
)

func GenerateJwtToken(secretKey string, iat, seconds, userID, state, isBoss int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims[ctxData.JwtKeyUserID] = userID
	claims[ctxData.JwtKeyUserState] = state
	claims[ctxData.JwtKeyIsBoss] = isBoss
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
