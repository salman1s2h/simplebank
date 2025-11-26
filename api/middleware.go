package api

import (
	// "go/token"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/salman1s2h/simplebank/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "authorization header is not provided"})
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization header format"})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "unsupported authorization type"})
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken, token.TokenTypeAccessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}

}
