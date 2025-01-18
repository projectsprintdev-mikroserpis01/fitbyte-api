package middlewares

import "github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/jwt"

type Middleware struct {
	jwt jwt.JwtInterface
}

func NewMiddleware(
	jwt jwt.JwtInterface,
) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}
